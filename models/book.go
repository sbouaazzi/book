// Author: Sami Bouaazzi
// Create Date: Apr 10, 2019

// Book object model for attributes of a Book.
//
// Defines and represents a book object and it's attributes.

package models

import "gopkg.in/mgo.v2/bson"

// Defines and represents a Book object and it's attributes, attribute types, and object formatting.
type Book struct {
	Id          bson.ObjectId `json:"id" bson:"_id"`                  //Id attribute to uniquely identify each book
	Title       string        `json:"title" bson:"title"`             //Title attribute that holds the title of the book
	Author      string        `json:"author" bson:"author"`           //Author attribute that holds the author of the book
	Publisher   string        `json:"publisher" bson:"publisher"`     //Publisher attribute that holds the publisher of the book
	PublishDate string        `json:"publishdate" bson:"publishdate"` //PublishDate attribute that holds the publish date of the book
	Rating      int           `json:"rating" bson:"rating"`           //Rating attribute that holds the rating number of the book
	Status      string        `json:"status" bson:"status"`           //Status attribute that holds the status of the book
}
