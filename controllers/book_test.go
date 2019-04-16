// Author: Sami Bouaazzi
// Create Date: Apr 10, 2019

// book.go Test Class
//
// Defines each book API route CRUD methods.
//
// Each method uses the instantiated 'dao' method to Create, Read, Update, and Delete data passed in from the route.
// The CRUD operations are operated on the MongoDB database.
// Each routed method returns a 200 or 400 response based on validation of the request and parameters, and whether it succeeds or fails.

package controllers

import (
	"bytes"
	"github.com/gorilla/mux"
	dao2 "github.com/sbouaazzi/book/dao"
	"github.com/sbouaazzi/book/models"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"net/http/httptest"
	"testing"
)

// test books reference for unit testing
var b = models.Book{Id: bson.NewObjectId(), Title: "Book Title", Author: "An Author", Publisher: "A Publisher", PublishDate: "0000", Rating: 1, Status: CheckedIn}
var b1 = models.Book{Id: bson.NewObjectId(), Title: "Book Title1", Author: "", Publisher: "A Publisher1", PublishDate: "1111", Rating: 1, Status: CheckedIn}
var b2 = models.Book{Id: bson.NewObjectId(), Title: "Book Title2", Author: "An Author2", Publisher: "A Publisher2", PublishDate: "2222", Rating: 5, Status: CheckedIn}
var b3 = models.Book{Id: bson.NewObjectId(), Title: "Book Title3", Author: "An Author3", Publisher: "A Publisher3", PublishDate: "3333", Rating: 5, Status: "ChEckEdiN"}
var b4 = models.Book{Id: bson.NewObjectId(), Title: "Book Title4", Author: "An Author4", Publisher: "A Publisher4", PublishDate: "12A45", Rating: 5, Status: CheckedIn}

// SetUp function
//
// Sets up test environment with stub book models inserted into the database.
func SetUp(db dao2.BookDAO) {
	db.Connect()
	_ = db.Insert(b)
	_ = db.Insert(b1)
	_ = db.Insert(b2)
	_ = db.Insert(b3)
	_ = db.Insert(b4)
}

// TearDown function
//
// Tears down the database environment deleting the test books from the database.
func TearDown(db dao2.BookDAO) {
	_ = db.Delete(b)
	_ = db.Delete(b1)
	_ = db.Delete(b2)
	_ = db.Delete(b3)
	_ = db.Delete(b4)
	db.Close()
}

// TestAllUnitTests function
//
// Run all unit tests from this function to stay in the same DB session
func TestAllUnitTests(t *testing.T) {
	// BookDAO instance with MongoDB URL and 'BookMongo' Database
	var db = dao2.BookDAO{Server: "mongodb://0.0.0.0:27017", Database: "BookMongo"}
	SetUp(db)
	GetAllBooksTest(t)
	GetBookTest(t)
	GetBadBookTest(t)
	CreateBookTest(t)
	CreateBadBookTest(t)
	UpdateBookTest(t)
	UpdateBadBookTest(t)
	DeleteBookTest(t)
	DeleteBadBookTest(t)
	RespondWithErrorTest(t)
	RespondWithJsonTest(t)
	TearDown(db)
}

// GetAllBooksTest function
// @param t - reference to testing.T object
//
// Runs unit test against the GetAllBooks method for a 200 http response with test http call
func GetAllBooksTest(t *testing.T) {

	req, err := http.NewRequest("GET", "/book", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllBooks)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

// GetBookTest function
// @param t - reference to testing.T object
//
// Runs unit test against the GetBook method for a 200 http response with test http call
func GetBookTest(t *testing.T) {
	req, err := http.NewRequest("GET", "/book/"+b.Id.Hex(), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetBook)

	vars := map[string]string{
		"id": b.Id.Hex(),
	}

	req = mux.SetURLVars(req, vars)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

// GetBadBookTest function
// @param t - reference to testing.T object
//
// Runs unit test against the GetBook method for a 400 http response with test http call
func GetBadBookTest(t *testing.T) {
	req, err := http.NewRequest("GET", "/book/5cae2195d67494bbc5532bba", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetBook)

	vars := map[string]string{
		"id": "5cae2195d67494bbc5532bba",
	}

	req = mux.SetURLVars(req, vars)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

}

// CreateBookTest function
// @param t - reference to testing.T object
//
// Runs unit test against the CreateBook method for a 200 http response with test http call
func CreateBookTest(t *testing.T) {
	payload := []byte(`{
	"title": "Gnat in the hat",
	"author": "Prof. Suess",
	"publisher": "Random Planet",
	"publishdate": "1969",
	"rating": 2,
	"status": "CheckedOut"
	}`)
	req, err := http.NewRequest("POST", "/book", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateBook)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

// CreateBadBookTest function
// @param t - reference to testing.T object
//
// Runs unit test against the CreateBadBook method for a 400 http response with test http call
func CreateBadBookTest(t *testing.T) {
	payload := []byte(`{
	"title": "",
	"author": "Prof. Suess",
	"publisher": "Random Planet",
	"publishdate": "1969",
	"rating": 2,
	"status": "CheckedOut"
	}`)
	req, err := http.NewRequest("POST", "/book", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateBook)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

}

// TestUpdateBook function
// @param t - reference to testing.T object
//
// Runs unit test against the UpdateBook method for a 200 http response with test http call
func UpdateBookTest(t *testing.T) {
	payload := []byte(`{
	"title": "Frat in the hat",
	"author": "Prof. Suess",
	"publisher": "Random Planet",
	"publishdate": "1969",
	"rating": 2,
	"status": "CheckedOut"
	}`)
	req, err := http.NewRequest("PUT", "/book/"+"/book/"+b.Id.Hex(), bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateBook)

	vars := map[string]string{
		"id": b.Id.Hex(),
	}

	req = mux.SetURLVars(req, vars)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

// UpdateBadBookTest function
// @param t - reference to testing.T object
//
// Runs unit test against the UpdateBook method for a 400 http response with test http call
func UpdateBadBookTest(t *testing.T) {
	payload := []byte(`{
	"title": "Frat in the hat",
	"author": "Prof. Suess",
	"publisher": "Random Planet",
	"publishdate": "1969",
	"rating": 40,
	"status": "CheckedOut"
	}`)
	req, err := http.NewRequest("PUT", "/book/"+"/book/"+b.Id.Hex(), bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateBook)

	vars := map[string]string{
		"id": b.Id.Hex(),
	}

	req = mux.SetURLVars(req, vars)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

}

// DeleteBookTest function
// @param t - reference to testing.T object
//
// Runs unit test against the DeleteBook method for a 200 http response with test http call
func DeleteBookTest(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/book/"+b.Id.Hex(), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteBook)

	vars := map[string]string{
		"id": b.Id.Hex(),
	}

	req = mux.SetURLVars(req, vars)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

// DeleteBadBookTest function
// @param t - reference to testing.T object
//
// Runs unit test against the DeleteBook method for a 400 http response with test http call
func DeleteBadBookTest(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/book/5cae2195d67494bbc5532bba", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteBook)

	vars := map[string]string{
		"id": "5cae2195d67494bbc5532bba",
	}

	req = mux.SetURLVars(req, vars)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

}

// Test_respondWithError function
// @param t - reference to testing.T object
//
// Runs unit test against the RespondWithError method with multiple scenarios
func RespondWithErrorTest(t *testing.T) {
	type args struct {
		w    http.ResponseWriter
		code int
		msg  string
	}
	tests := []struct {
		name string
		args args
	}{
		// Test cases for validating JSON error responses
		{
			name: ValidTest1,
			args: args{
				w:    httptest.NewRecorder(),
				code: 400,
				msg:  "System Error",
			},
		},
		{
			name: ValidTest2,
			args: args{
				w:    httptest.NewRecorder(),
				code: 400,
				msg:  "Error",
			},
		},
		{
			name: Not + ValidTest1,
			args: args{
				w:    httptest.NewRecorder(),
				code: 400,
				msg:  "",
			},
		},
		{
			name: Not + ValidTest2,
			args: args{
				w:    httptest.NewRecorder(),
				code: 400,
				msg: "		",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			respondWithError(tt.args.w, tt.args.code, tt.args.msg)
		})
	}
}

// Test_respondWithJson function
// @param t - reference to testing.T object
//
// Runs unit test against the RespondWithJson method with multiple scenarios
func RespondWithJsonTest(t *testing.T) {
	type args struct {
		w       http.ResponseWriter
		code    int
		payload interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// Test cases for validating JSON responses
		{
			name: ValidTest1,
			args: args{
				w:       httptest.NewRecorder(),
				code:    200,
				payload: map[string]string{Result: Success},
			},
		},
		{
			name: ValidTest2,
			args: args{
				w:       httptest.NewRecorder(),
				code:    200,
				payload: map[string]string{Error: "error"},
			},
		},
		{
			name: Not + ValidTest1,
			args: args{
				w:       httptest.NewRecorder(),
				code:    200,
				payload: "Good",
			},
		},
		{
			name: Not + ValidTest2,
			args: args{
				w:       httptest.NewRecorder(),
				code:    400,
				payload: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			respondWithJson(tt.args.w, tt.args.code, tt.args.payload)
		})
	}
}
