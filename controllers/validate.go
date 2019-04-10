package controllers

import (
	"book/models"
	"github.com/asaskevich/govalidator"
)

const EMPTY_STRING = ""

var (
	// ErrInvalidRange - error when we have a rating range validation issue
	ErrInvalidRatingRange = "Error: invalid rating range"
	// ErrInvalidEntry - error when we have an invalid entry on any field
	ErrInvalidEntry = "Error: invalid text entry"
	// ErrInvalidDate - error when we have an invalid publish date entry
	ErrInvalidDate = "Error: invalid date entry"
	// ErrInvalidStatus - error when we have an invalid status entry
	ErrInvalidStatus = "Error: invalid status entry"
)

// Validate - implementation of the InputValidation interface
func Validate(b models.Book) string {
	// validate the rating range is between 1 and 3
	if RatingValidator(b.Rating) {
		return ErrInvalidRatingRange
	}

	// validate string entries are not empty, missing, or invalid (only whitespaces)
	values := []string{b.Title, b.Author, b.Publisher}
	for _, s := range values {
		// validate the entry is not empty or missing
		if StringEntryValidator(s) {
			return ErrInvalidEntry
		}
	}

	// validate the status entry
	if StatusValidator(b.Status) {
		return ErrInvalidStatus
	}

	// validate the publish date entry
	if PublishDateValidator(b.PublishDate) {
		return ErrInvalidDate
	}

	return EMPTY_STRING
}

func StringEntryValidator(s string) bool {
	// validate the entry is not empty or missing
	if govalidator.IsNull(s) || s == EMPTY_STRING || govalidator.HasWhitespaceOnly(s) {
		return true
	}

	return false
}

func StatusValidator(s string) bool {
	// validate the entry is not empty or missing
	if StringEntryValidator(s) || (s != "CheckedIn" && s != "CheckedOut") {
		return true
	}

	return false
}

func RatingValidator(i int) bool {
	if i > 3 || i < 1 {
		return true
	}

	return false
}

func PublishDateValidator(s string) bool {
	if StringEntryValidator(s) || !govalidator.IsNumeric(s) {
		return true
	}

	return false
}
