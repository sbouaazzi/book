package main

import (
	"book/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
)

const (
	ApplicationPortNumberMsg = ":8080"
	ApplicationStartedMsg    = "Application Started"
	BookEndpoint             = "/book"
	BookIdParamEndpoint      = "/book/:id"
	ListeningOnPort8080Msg   = "Listening on port 8080"
	ContentType              = "Content-Type"
	DatabaseServerUrl        = "mongodb://book_mongodb_1:27017"
	ErrorGettingSessionMsg   = "Error getting session: %s\n"
)

func main() {
	log.Println(ApplicationStartedMsg)
	r := httprouter.New()
	bc := controllers.NewBookController(getSession())

	//GET
	r.GET(BookEndpoint, bc.GetAllBooks)
	r.GET(BookIdParamEndpoint, bc.GetBook)

	//POST
	r.POST(BookEndpoint, bc.CreateBook)

	//DELETE
	r.DELETE(BookIdParamEndpoint, bc.DeleteBook)

	//PUT
	r.PUT(BookIdParamEndpoint, bc.UpdateBook)

	log.Println(ListeningOnPort8080Msg)
	_ = http.ListenAndServe(ApplicationPortNumberMsg, r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial(DatabaseServerUrl)

	if err != nil {
		log.Panicf(ErrorGettingSessionMsg, err.Error())
	}
	return s
}
