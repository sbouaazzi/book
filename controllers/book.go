package controllers

import (
	"book/models"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

//The BookController struct containing the session of the model Book
type BookController struct {
	session *mgo.Session
}

//The NewBookController function
func NewBookController(s *mgo.Session) *BookController {
	return &BookController{s}
}

//The GetBook function
func (bc BookController) GetBook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	oid := bson.ObjectIdHex(id)

	b := models.Book{}

	if err := bc.session.DB("BookMongo").C("books").FindId(oid).One(&b); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	bj, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", bj)
}

//The CreateBook function
func (bc BookController) CreateBook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	b := models.Book{}

	json.NewDecoder(r.Body).Decode(&b)

	errorMsg := Validate(b)
	if errorMsg != EMPTY_STRING {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "%s\n", errorMsg)
		return
	}

	b.Id = bson.NewObjectId()

	bc.session.DB("BookMongo").C("books").Insert(b)

	bj, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", bj)
}

//The UpdateBook function
func (bc BookController) UpdateBook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	b := models.Book{}
	id := p.ByName("id")
	oid := bson.ObjectIdHex(id)

	json.NewDecoder(r.Body).Decode(&b)

	errorMsg := Validate(b)
	if errorMsg != EMPTY_STRING {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "%s\n", errorMsg)
		return
	}

	b.Id = bson.ObjectIdHex(id)

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	bc.session.DB("BookMongo").C("books").UpdateId(oid, &b)

	bj, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 201
	fmt.Fprintf(w, "%s\n", bj)
}

//The DeleteBook Function
func (bc BookController) DeleteBook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(id)

	// Delete user
	if err := bc.session.DB("BookMongo").C("books").RemoveId(oid); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user ", oid, "\n")
}
