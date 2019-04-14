// Author: Sami Bouaazzi
// Create Date: Apr 12, 2019

// Validates the Book's attributes for valid entries and inputs.
//
// The main Validate method splits up into 4 separate methods in order to validate the different attribute data types and rules.
// The validations include:
//			-Rating number range must be 1-3
//			-Status values equal either "CheckedIn" or "CheckedOut"
//			-Title, Author, and Publisher are not blank or empty values
//			-Publisher Date is not blank or empty and a numeric value

package controllers

import (
	"book/models"
	"github.com/asaskevich/govalidator"
)

// constants definitions
const (
	CheckedIn             = "CheckedIn"
	CheckedOut            = "CheckedOut"
	ErrInvalidDate        = "Invalid publish date entry. Value must be a numeric year, and not be blank or empty."
	ErrInvalidRatingRange = "Invalid rating range. Value range must be from 1-3."
	ErrInvalidStatus      = "Invalid status entry. Value must be either 'CheckedIn' or 'CheckedOut', and not be blank or empty."
	ErrInvalidTextEntry   = "Invalid text entry. Values must not be blank or empty."
)

// Validate function
// @param b - the Book object
//
// Validates a book record from the given parameter.
// The validations include:
//			-Rating value range must be 1-3
//			-Status value equals either "CheckedIn" or "CheckedOut" and not blank or empty values
//			-Title, Author, and Publisher values are not blank or empty values
//			-Publisher Date value is not blank or empty and a numeric value
//
// The function returns a string with the validation error message, or an empty string if none are found
func Validate(b models.Book) string {
	// validate string entries are not empty, missing, or invalid (only whitespaces)
	values := []string{b.Title, b.Author, b.Publisher}
	for _, s := range values {
		// validate the entry is not empty or missing
		if StringEntryValidator(s) {
			return ErrInvalidTextEntry
		}
	}

	// validate the rating range is between 1 and 3
	if RatingValidator(b.Rating) {
		return ErrInvalidRatingRange
	}

	// validate the status entry
	if StatusValidator(b.Status) {
		return ErrInvalidStatus
	}

	// validate the publish date entry
	if PublishDateValidator(b.PublishDate) {
		return ErrInvalidDate
	}

	return EmptyString
}

// StringEntryValidator function
// @param s - the Book string attribute value
//
// Validates a book's string attribute value from the given parameter.
// The validation includes that the values are not blank, null or empty values
// and not a blank, null or empty value.
// The function returns a boolean. True if a validation error is found, false if not.
func StringEntryValidator(s string) bool {
	// validate the entry is not empty or missing
	if govalidator.IsNull(s) || s == EmptyString || govalidator.HasWhitespaceOnly(s) {
		return true
	}

	return false
}

// StatusValidator function
// @param s - the Book.Status string value
//
// Validates a book's status value from the given parameter.
// The validation includes that the status value equals either "CheckedIn" or "CheckedOut"
// and not blank, null or empty value.
// The function returns a boolean. True if a validation error is found, false if not.
func StatusValidator(s string) bool {
	// validate the entry is not empty or missing
	if StringEntryValidator(s) || (s != CheckedIn && s != CheckedOut) {
		return true
	}

	return false
}

// RatingValidator function
// @param i - the Book.Rating int value
//
// Validates a book's rating value from the given parameter.
// The validation includes that the rating number range must be 1-3.
// The function returns a boolean. True if a validation error is found, false if not.
func RatingValidator(i int) bool {
	if i > 3 || i < 1 {
		return true
	}

	return false
}

// PublishDateValidator function
// @param s - the Book.PublishDate string value
//
// Validates a book's publish date value from the given parameter.
// The validation includes that the publish date is not blank, null or empty values.
// The function returns a boolean. True if a validation error is found, false if not.
func PublishDateValidator(s string) bool {
	if StringEntryValidator(s) || !govalidator.IsNumeric(s) {
		return true
	}

	return false
}
