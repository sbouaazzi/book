// Author: Sami Bouaazzi
// Create Date: Apr 14, 2019

// Validate.go Test Class
//
// Validates the Book's attributes for valid entries and inputs.
//
// The main Validate method splits up into 4 separate methods in order to validate the different attribute data types and rules.
// The validations include:
//			-Rating value range must be 1-3
//			-Status value equals either "CheckedIn" or "CheckedOut" and not blank or empty values with case sensitivity
//			-Title, Author, and Publisher values are not blank or empty values
//			-Publisher Date value is not blank or empty and a numeric value

package controllers

import (
	"github.com/sbouaazzi/book/models"
	"testing"
)

// constants definitions
const (
	Not        = "Not"
	ValidTest1 = "ValidTest1"
	ValidTest2 = "ValidTest2"
	ValidTest3 = "ValidTest3"
	ValidTest4 = "ValidTest4"
	ValidTest5 = "ValidTest5"
	ValidTest6 = "ValidTest6"
	ValidTest7 = "ValidTest7"
)

// TestValidate function
// @param t - reference to testing.T object
//
// Runs test cases against the Validate method with valid and invalid scenarios for a Book object attributes
func TestValidate(t *testing.T) {
	type args struct {
		b models.Book
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// Test cases for validating book objects
		{
			name: ValidTest1,
			args: args{
				models.Book{Id: "12345", Title: "Book Title", Author: "An Author", Publisher: "A Publisher", PublishDate: "1234", Rating: 1, Status: CheckedIn},
			},
			want: EmptyString,
		},
		{
			name: ValidTest2,
			args: args{
				models.Book{Id: "12345", Title: "Book Title", Author: "An Author", Publisher: "A Publisher", PublishDate: "1234", Rating: 3, Status: CheckedOut},
			},
			want: EmptyString,
		},
		{
			name: Not + ValidTest1,
			args: args{
				models.Book{Id: "12345", Title: "Book Title", Author: "An Author", Publisher: "A Publisher", PublishDate: "ABCD", Rating: 3, Status: CheckedOut},
			},
			want: ErrInvalidDate,
		},
		{
			name: Not + ValidTest2,
			args: args{
				models.Book{Id: "12345", Title: "Book Title", Author: "An Author", Publisher: "A Publisher", PublishDate: "1234", Status: CheckedIn},
			},
			want: ErrInvalidRatingRange,
		},
		{
			name: Not + ValidTest3,
			args: args{
				models.Book{Id: "12345", Title: "Book Title", Author: "An Author", Publisher: "A Publisher", PublishDate: "1234", Rating: 2, Status: "checkedin"},
			},
			want: ErrInvalidStatus,
		},
		{
			name: "NotValidTest4",
			args: args{
				models.Book{Id: "12345", Author: "An Author", PublishDate: "1234", Rating: 1, Status: CheckedIn},
			},
			want: ErrInvalidTextEntry,
		},
		{
			name: Not + ValidTest5,
			args: args{
				models.Book{Status: "CHeCkEDouT"},
			},
			want: ErrInvalidTextEntry,
		},
		{
			name: Not + ValidTest6,
			args: args{
				models.Book{Id: "12345", Title: "Book Title", Author: "An Author", Publisher: "A Publisher", PublishDate: "ABCDE", Status: "CHeCkEDouT"},
			},
			want: ErrInvalidRatingRange,
		},
		{
			name: Not + ValidTest7,
			args: args{
				models.Book{Id: "12345", Title: "Book Title", Author: "An Author", Publisher: "A Publisher", PublishDate: "ABCDE", Rating: 3, Status: "CHeCkEDiN"},
			},
			want: ErrInvalidStatus,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Validate(tt.args.b); got != tt.want {
				t.Errorf("Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestStringEntryValidator function
// @param t - reference to testing.T object
//
// Runs test cases against the StringEntryValidator method with valid and invalid scenarios for a string value
func TestStringEntryValidator(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// Test cases for validating string values
		{
			name: ValidTest1,
			args: args{
				"A Text Value",
			},
			want: false,
		},
		{
			name: ValidTest2,
			args: args{
				"12345",
			},
			want: false,
		},
		{
			name: ValidTest3,
			args: args{
				"ABCDE",
			},
			want: false,
		},
		{
			name: ValidTest4,
			args: args{
				"@@@##*",
			},
			want: false,
		},
		{
			name: Not + ValidTest1,
			args: args{
				"",
			},
			want: true,
		},
		{
			name: Not + ValidTest2,
			args: args{
				"	",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringEntryValidator(tt.args.s); got != tt.want {
				t.Errorf("StringEntryValidator() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestStatusValidator function
// @param t - reference to testing.T object
//
// Runs test cases against the StatusValidator method with valid and invalid scenarios for a Book.Status value
func TestStatusValidator(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// Test cases for validating status string values
		{
			name: ValidTest1,
			args: args{
				CheckedIn,
			},
			want: false,
		},
		{
			name: ValidTest2,
			args: args{
				CheckedOut,
			},
			want: false,
		},
		{
			name: Not + ValidTest1,
			args: args{
				"checkedin",
			},
			want: true,
		},
		{
			name: Not + ValidTest2,
			args: args{
				"checkedout",
			},
			want: true,
		},
		{
			name: Not + ValidTest3,
			args: args{
				"CHECKEDIN",
			},
			want: true,
		},
		{
			name: Not + ValidTest4,
			args: args{
				"CHECKEDOUT",
			},
			want: true,
		},
		{
			name: Not + ValidTest5,
			args: args{
				"   ",
			},
			want: true,
		},
		{
			name: Not + ValidTest6,
			args: args{
				"",
			},
			want: true,
		},
		{
			name: Not + ValidTest7,
			args: args{
				"12345",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StatusValidator(tt.args.s); got != tt.want {
				t.Errorf("StatusValidator() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestRatingValidator function
// @param t - reference to testing.T object
//
// Runs test cases against the RatingValidator method with valid and invalid scenarios for a Book.Rating value
func TestRatingValidator(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// Test cases for validating rating number values
		{
			name: ValidTest1,
			args: args{
				1,
			},
			want: false,
		},
		{
			name: ValidTest2,
			args: args{
				2,
			},
			want: false,
		},
		{
			name: ValidTest3,
			args: args{
				3,
			},
			want: false,
		},
		{
			name: Not + ValidTest1,
			args: args{
				0,
			},
			want: true,
		},
		{
			name: Not + ValidTest2,
			args: args{
				4,
			},
			want: true,
		},
		{
			name: Not + ValidTest3,
			args: args{
				-3,
			},
			want: true,
		},
		{
			name: Not + ValidTest4,
			args: args{
				100,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RatingValidator(tt.args.i); got != tt.want {
				t.Errorf("RatingValidator() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestPublishDateValidator function
// @param t - reference to testing.T object
//
// Runs test cases against the PublishDateValidator method with valid and invalid scenarios for a Book.PublishDate value
func TestPublishDateValidator(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// Test cases for validating publish date string values
		{
			name: ValidTest1,
			args: args{
				"1960",
			},
			want: false,
		},
		{
			name: ValidTest2,
			args: args{
				"2019",
			},
			want: false,
		},
		{
			name: Not + ValidTest1,
			args: args{
				"ABCD",
			},
			want: true,
		},
		{
			name: Not + ValidTest2,
			args: args{
				"2015A",
			},
			want: true,
		},
		{
			name: Not + ValidTest3,
			args: args{
				"20152",
			},
			want: true,
		},
		{
			name: Not + ValidTest4,
			args: args{
				"",
			},
			want: true,
		},
		{
			name: Not + ValidTest5,
			args: args{
				"   ",
			},
			want: true,
		},
		{
			name: Not + ValidTest6,
			args: args{
				"!!@@",
			},
			want: true,
		},
		{
			name: Not + ValidTest7,
			args: args{
				"201",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PublishDateValidator(tt.args.s); got != tt.want {
				t.Errorf("PublishDateValidator() = %v, want %v", got, tt.want)
			}
		})
	}
}
