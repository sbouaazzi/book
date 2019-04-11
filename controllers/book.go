package controllers

import (
	"book/models"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

const (
	ApplicationJSON  = "application/json"
	CollectionName   = "books"
	ContentType      = "Content-Type"
	DatabaseName     = "BookMongo"
	DeletedIdMessage = "Deleted book id "
	Error            = "Error: "
	EmptyString      = ""
	Format           = "%s\n"
	ID               = "id"
	NewLine          = "\n"
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
func (bc BookController) GetAllBooks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var b []models.Book

	if err := bc.session.DB(DatabaseName).C(CollectionName).Find(nil).All(&b); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	bj, err := json.Marshal(b)
	if err != nil {
		log.Panic(Error, err.Error())
	}

	w.Header().Set(ContentType, ApplicationJSON)
	w.WriteHeader(http.StatusOK) // 200
	_, _ = fmt.Fprintf(w, Format, bj)
}

//The GetBook function
func (bc BookController) GetBook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName(ID)

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	oid := bson.ObjectIdHex(id)

	b := models.Book{}

	if err := bc.session.DB(DatabaseName).C(CollectionName).FindId(oid).One(&b); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	bj, err := json.Marshal(b)
	if err != nil {
		log.Panic(Error, err.Error())
	}

	w.Header().Set(ContentType, ApplicationJSON)
	w.WriteHeader(http.StatusOK) // 200
	_, _ = fmt.Fprintf(w, Format, bj)
}

//The CreateBook function
func (bc BookController) CreateBook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	b := models.Book{}

	_ = json.NewDecoder(r.Body).Decode(&b)

	errorMsg := Validate(b)
	if errorMsg != EmptyString {
		w.Header().Set(ContentType, ApplicationJSON)
		w.WriteHeader(http.StatusBadRequest) // 400
		_, _ = fmt.Fprintf(w, Format, errorMsg)
		return
	}

	b.Id = bson.NewObjectId()

	_ = bc.session.DB(DatabaseName).C(CollectionName).Insert(b)

	bj, err := json.Marshal(b)
	if err != nil {
		log.Panic(Error, err.Error())
	}

	w.Header().Set(ContentType, ApplicationJSON)
	w.WriteHeader(http.StatusOK) // 200
	_, _ = fmt.Fprintf(w, Format, bj)
}

//The UpdateBook function
func (bc BookController) UpdateBook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	b := models.Book{}
	id := p.ByName(ID)
	oid := bson.ObjectIdHex(id)

	_ = json.NewDecoder(r.Body).Decode(&b)

	errorMsg := Validate(b)
	if errorMsg != EmptyString {
		w.Header().Set(ContentType, ApplicationJSON)
		w.WriteHeader(http.StatusBadRequest) // 400
		_, _ = fmt.Fprintf(w, Format, errorMsg)
		return
	}

	b.Id = bson.ObjectIdHex(id)

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	_ = bc.session.DB(DatabaseName).C(CollectionName).UpdateId(oid, &b)

	bj, err := json.Marshal(b)
	if err != nil {
		log.Panic(Error, err.Error())
	}

	w.Header().Set(ContentType, ApplicationJSON)
	w.WriteHeader(http.StatusOK) // 200
	_, _ = fmt.Fprintf(w, Format, bj)
}

//The DeleteBook Function
func (bc BookController) DeleteBook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName(ID)

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(id)

	// Delete user
	if err := bc.session.DB(DatabaseName).C(CollectionName).RemoveId(oid); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK) // 200
	_, _ = fmt.Fprint(w, DeletedIdMessage, oid, NewLine)
}
