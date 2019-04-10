package models

import "gopkg.in/mgo.v2/bson"

type Book struct {
	Id          bson.ObjectId `json:"id" bson:"_id"`
	Title       string        `json:"title" bson:"title"`
	Author      string        `json:"author" bson:"author"`
	Publisher   string        `json:"publisher" bson:"publisher"`
	PublishDate string        `json:"publishdate" bson:"publishdate"`
	Rating      int           `json:"rating" bson:"rating"`
	Status      string        `json:"status" bson:"status"`
}
