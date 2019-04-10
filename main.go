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
	r.GET("/book/:id", bc.GetBook)
	r.POST("/book", bc.CreateBook)
	r.DELETE("/book/:id", bc.DeleteBook)
	r.PUT("/book/:id", bc.UpdateBook)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}
	return s
}
