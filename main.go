package main

import (
	"book/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"net/http"
)

func main() {
	r := httprouter.New()
	bc := controllers.NewBookController(getSession())
	//GET
	r.GET("/book", bc.GetAllBooks)
	r.GET("/book/:id", bc.GetBook)

	//POST
	r.POST("/book", bc.CreateBook)

	//DELETE
	r.DELETE("/book/:id", bc.DeleteBook)

	//PUT
	r.PUT("/book/:id", bc.UpdateBook)
	http.ListenAndServe(":8080", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://book_mongodb_1:27017")

	if err != nil {
		panic(err)
	}
	return s
}
