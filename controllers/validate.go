package controllers

import (
	"book/models"
	"github.com/asaskevich/govalidator"
)

const (
	CheckedIn             = "CheckedIn"
	CheckedOut            = "CheckedOut"
	BlankString           = ""
	ErrInvalidDate        = "Error: invalid date entry"
	ErrInvalidEntry       = "Error: invalid text entry"
	ErrInvalidRatingRange = "Error: invalid rating range"
	ErrInvalidStatus      = "Error: invalid status entry"
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

	return BlankString
}

func StringEntryValidator(s string) bool {
	// validate the entry is not empty or missing
	if govalidator.IsNull(s) || s == BlankString || govalidator.HasWhitespaceOnly(s) {
		return true
	}

	return false
}

func StatusValidator(s string) bool {
	// validate the entry is not empty or missing
	if StringEntryValidator(s) || (s != CheckedIn && s != CheckedOut) {
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
