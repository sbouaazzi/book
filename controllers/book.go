// Author: Sami Bouaazzi
// Create Date: Apr 10, 2019

// Defines each book API route CRUD methods.
//
// Each method uses the instantiated 'dao' method to Create, Read, Update, and Delete data passed in from the route.
// The CRUD operations are operated on the MongoDB database.
// Each routed method returns a 200 or 400 response based on validation of the request and parameters, and whether it succeeds or fails.

package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	dao2 "github.com/sbouaazzi/book/dao"
	"github.com/sbouaazzi/book/models"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

// constants definitions
const (
	ApplicationJSON          = "application/json"
	ContentType              = "Content-Type"
	CreatedBookMsg           = "Created book with record id: "
	DatabaseName             = "BookMongo"
	DatabaseServerUrl        = "mongodb://book_mongodb_1:27017"
	DeletedBookMsg           = "Deleted book with record id: "
	Error                    = "Error: "
	EmptyString              = ""
	Format                   = "%s\n"
	InvalidBookIdMsg         = "Invalid book id"
	InvalidRequestPayloadMsg = "Invalid request payload"
	Id                       = "id"
	Result                   = "Result"
	RetrievedAllBooksMsg     = "Retrieved all books"
	RetrievedBookMsg         = "Retrieved book with record id: "
	Success                  = "Success"
	UpdatedBookMsg           = "Updated book with record id: "
)

// BookDAO instance with MongoDB URL and 'BookMongo' Database
var dao = dao2.BookDAO{Server: DatabaseServerUrl, Database: DatabaseName}

// GetAllBooks function
// @param w - the ResponseWriter object
// @param r - pointer to the Request object
//
// HTTP GET '/book' route method
// Retrieves all book records from the 'BookMongo' Database.
// The function responds with a 200 status and the JSON list of book records with id's for successful retrieval.
// The function responds with a 400 status and a JSON error message for any system errors.
func GetAllBooks(w http.ResponseWriter, _ *http.Request) {
	books, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	log.Println(RetrievedAllBooksMsg)

	respondWithJson(w, http.StatusOK, books)
}

// GetBook function
// @param w - the ResponseWriter object
// @param r - pointer to the Request object
//
// HTTP GET '/book/{id}' route method with an id parameter.
// Retrieves one book record from the 'BookMongo' Database given a valid {id} parameter.
// The function validates whether the entered id exists in the database.
// The function responds with a 200 status and a JSON message of the specified book record for a successful retrieval.
// The function responds with a 400 status and a JSON error message for any system errors or validation errors.
func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	book, err := dao.FindById(params[Id])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, InvalidBookIdMsg)
		return
	}
	log.Println(RetrievedBookMsg + book.Id.Hex())

	respondWithJson(w, http.StatusOK, book)
}

// CreateBook function
// @param w - the ResponseWriter object
// @param r - pointer to the Request object
//
// HTTP POST '/book' route method with a JSON request message.
// Creates one book record for the 'BookMongo' Database given valid JSON request message formatting.
// The function validates the each JSON field for correct formatting. (no empty values or strings, rating is between 1-3, Status equals 'CheckedIn' or 'CheckedOut', etc).
// The function responds with a 200 status and a JSON message of the created book record for a successful creation.
// The function responds with a 400 status and a JSON error message for any system errors or validation errors.
func CreateBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		respondWithError(w, http.StatusBadRequest, InvalidRequestPayloadMsg)
		return
	}

	errorMsg := Validate(book)
	if errorMsg != EmptyString {
		respondWithError(w, http.StatusBadRequest, errorMsg)
		return
	}

	book.Id = bson.NewObjectId()

	if err := dao.Insert(book); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	log.Println(CreatedBookMsg + book.Id.Hex())

	respondWithJson(w, http.StatusOK, book)
}

// UpdateBook function
// @param w - the ResponseWriter object
// @param r - pointer to the Request object
//
// HTTP PUT '/book/{id}' route method with an id parameter and a JSON request message.
// Updates one book record on the 'BookMongo' Database given a valid {id} parameter and valid JSON request message formatting.
// The function validates the each JSON field for correct formatting. (no empty values or strings, rating is between 1-3, Status equals 'CheckedIn' or 'CheckedOut', etc).
// The function also validates whether the entered id exists in the database.
// The function responds with a 200 status and a JSON message of the updated book record for a successful update.
// The function responds with a 400 status and a JSON error message for any system errors or validation errors.
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	oldbook, _ := dao.FindById(params[Id])

	defer r.Body.Close()
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		respondWithError(w, http.StatusBadRequest, InvalidRequestPayloadMsg)
		return
	}

	book.Id = oldbook.Id

	errorMsg := Validate(book)
	if errorMsg != EmptyString {
		respondWithError(w, http.StatusBadRequest, errorMsg)
		return
	}

	if err := dao.Update(book); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	log.Println(UpdatedBookMsg + book.Id.Hex())

	respondWithJson(w, http.StatusOK, book)
}

// DeleteBook function
// @param w - the ResponseWriter object
// @param r - pointer to the Request object
//
// HTTP DELETE '/book/{id}' route method with an id parameter.
// Deletes one book record from the 'BookMongo' Database given a valid {id} parameter.
// The function validates whether the entered id exists in the database.
// The function responds with a 200 status and a 'Result: Success' JSON message for a successful deletion.
// The function responds with a 400 status and a JSON error message for any system errors or validation errors.
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	book, err := dao.FindById(params[Id])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, InvalidBookIdMsg)
		return
	}

	if err := dao.Delete(book); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	log.Println(DeletedBookMsg + book.Id.Hex())
	respondWithJson(w, http.StatusOK, map[string]string{Result: Success})
}

// respondWithJson function
// @param w - the ResponseWriter object
// @param code - the HTTP status code
// @param payload - a payload interface that takes in map[string]string
//
// The function takes the parameter arguments, inserts the error message into a map with Key "error",
// and calls the respondWithJson function to return JSON message.
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{Error: msg})
}

// respondWithJson function
// @param w - the ResponseWriter object
// @param code - the HTTP status code
// @param payload - a payload interface that takes in map[string]string
//
// The function takes the parameter arguments, and writes into the ResponseWriter the JSON message to return.
// with the given payload argument.
func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set(ContentType, ApplicationJSON)
	w.WriteHeader(code)
	_, _ = w.Write(response)
}
