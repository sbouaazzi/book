// Author: Sami Bouaazzi
// Create Date: Apr 11, 2019

// Data Access Object (DAO) class implementation to MongoDB.
//
// Handles the MongoDB database CRUD operations such as Find, FindId, Insert, Remove, and UpdateId.
// The DAO is also responsible for connecting to the MongoDB session, or Closing the session with the
// Connect() and Close() methods.

package dao

import (
	"book/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

// A BookDAO represents the information needed to access MongoDB
type BookDAO struct {
	Server   string
	Database string
}

// db variable referencing the mgo.Database type
var db *mgo.Database

// constants definitions
const (
	ClosedMongoSessionMsg = "Closed connection to MongoDB"
	Collection            = "books"
	ConnectedToMongoMsg   = "Connected to MongoDB"
)

// Connect function
//
// The function establishes a connection session to MongoDB for the DAO.
func (b *BookDAO) Connect() {
	session, err := mgo.Dial(b.Server)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(ConnectedToMongoMsg)
	db = session.DB(b.Database)
}

// Close function
//
// The function closes the MongoDB session for the DAO.
func (b *BookDAO) Close() {
	db.Session.Close()
	log.Println(ClosedMongoSessionMsg)
}

// FindAll function
//
// The function finds a list of all books from the collection.
// The functions returns an array of book objects and an error, which will be nil if the books are retrieved.
func (b *BookDAO) FindAll() ([]models.Book, error) {
	var books []models.Book
	err := db.C(Collection).Find(nil).All(&books)
	return books, err
}

// FindById function
// @param id - the id value
//
// The function finds a book by its id from the collection for a given id parameter.
// The functions returns a book object and an error, which will be nil if the book is retrieved.
func (b *BookDAO) FindById(id string) (models.Book, error) {
	var book models.Book
	err := db.C(Collection).FindId(bson.ObjectIdHex(id)).One(&book)
	return book, err
}

// Insert function
// @param book - a Book object
//
// The function inserts a new book from the collection for a given Book parameter.
// The functions returns error, which will be nil for a successful transaction or an
// error for any encountered errors.
func (b *BookDAO) Insert(book models.Book) error {
	err := db.C(Collection).Insert(&book)
	return err
}

// Update function
// @param book - a Book object
//
// The function deletes an existing book from the collection for a given Book parameter.
// The functions returns error, which will be nil for a successful transaction or an
// error for any encountered errors.
func (b *BookDAO) Delete(book models.Book) error {
	err := db.C(Collection).Remove(&book)
	return err
}

// Update function
// @param book - a Book object
//
// The function updates an existing book from the collection for a given Book parameter.
// The functions returns error, which will be nil for a successful transaction or an
// error for any encountered errors.
func (b *BookDAO) Update(book models.Book) error {
	err := db.C(Collection).UpdateId(book.Id, &book)
	return err
}
