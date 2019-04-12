package dao

import (
	"book/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type BookDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "books"
)

// Establish a connection to database
func (b *BookDAO) Connect() {
	session, err := mgo.Dial(b.Server)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")
	db = session.DB(b.Database)
}

// Close a connection to database
func (b *BookDAO) Close() {
	db.Session.Close()
}

// Find list of books
func (b *BookDAO) FindAll() ([]models.Book, error) {
	var books []models.Book
	err := db.C(COLLECTION).Find(nil).All(&books)
	return books, err
}

// Find a book by its id
func (b *BookDAO) FindById(id string) (models.Book, error) {
	var book models.Book
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&book)
	return book, err
}

// Insert a book into database
func (b *BookDAO) Insert(book models.Book) error {
	err := db.C(COLLECTION).Insert(&book)
	return err
}

// Delete an existing book
func (b *BookDAO) Delete(book models.Book) error {
	err := db.C(COLLECTION).Remove(&book)
	return err
}

// Update an existing book
func (b *BookDAO) Update(book models.Book) error {
	err := db.C(COLLECTION).UpdateId(book.Id, &book)
	return err
}
