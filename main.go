package main

import (
	"book/controllers"
	dao2 "book/dao"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	ApplicationPortNumberMsg = ":8080"
	ApplicationStartedMsg    = "Application Started"
	BookEndpoint             = "/book"
	BookIdParamEndpoint      = "/book/{id}"
	ListeningOnPort8080Msg   = "Listening on port 8080"
	ContentType              = "Content-Type"
)

var dao = dao2.BookDAO{Server: "mongodb://book_mongodb_1:27017", Database: "BookMongo"}

func main() {
	log.Println(ApplicationStartedMsg)
	dao.Connect()
	r := mux.NewRouter()

	//GET
	r.HandleFunc(BookEndpoint, controllers.GetAllBooks).Methods(http.MethodGet)
	r.HandleFunc(BookIdParamEndpoint, controllers.GetBook).Methods(http.MethodGet)

	//POST
	r.HandleFunc(BookEndpoint, controllers.CreateBook).Methods(http.MethodPost)

	//PUT
	r.HandleFunc(BookIdParamEndpoint, controllers.UpdateBook).Methods(http.MethodPut)

	//DELETE
	r.HandleFunc(BookIdParamEndpoint, controllers.DeleteBook).Methods(http.MethodDelete)

	log.Println(ListeningOnPort8080Msg)
	if err := http.ListenAndServe(ApplicationPortNumberMsg, r); err != nil {
		log.Fatal(err)
	}
}
