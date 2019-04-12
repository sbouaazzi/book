package controllers

import (
	dao2 "book/dao"
	"book/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

const (
	ApplicationJSON   = "application/json"
	ContentType       = "Content-Type"
	DatabaseName      = "BookMongo"
	DatabaseServerUrl = "mongodb://book_mongodb_1:27017"
	Error             = "Error: "
	EmptyString       = ""
	Format            = "%s\n"
	ID                = "id"
	Result            = "Result"
	Success           = "Success"
)

var dao = dao2.BookDAO{Server: DatabaseServerUrl, Database: DatabaseName}

//The GetBook function
func GetAllBooks(w http.ResponseWriter, _ *http.Request) {
	books, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	log.Println("Retrieved all books")

	respondWithJson(w, http.StatusOK, books)
}

//The GetBook function
func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	book, err := dao.FindById(params["id"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Book ID")
		return
	}
	log.Println("Retrieved book with record id: " + book.Id)

	respondWithJson(w, http.StatusOK, book)
}

//The CreateBook function
func CreateBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
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
	log.Println("Created book with record id: " + book.Id)

	respondWithJson(w, http.StatusOK, book)
}

//The UpdateBook function
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	oldbook, _ := dao.FindById(params["id"])

	defer r.Body.Close()
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
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
	log.Println("Updated book with record id: " + book.Id)

	respondWithJson(w, http.StatusOK, book)
}

//The DeleteBook Function
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	book, err := dao.FindById(params["id"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Book ID")
		return
	}

	if err := dao.Delete(book); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	log.Println("Deleted book with record id: " + book.Id)
	respondWithJson(w, http.StatusOK, map[string]string{Result: Success})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set(ContentType, ApplicationJSON)
	w.WriteHeader(code)
	_, _ = w.Write(response)
}
