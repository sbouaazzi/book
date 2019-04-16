// Author: Sami Bouaazzi
// Create Date: Apr 10, 2019

// Main application file to define routes and run application.
//
// Starts up the book api application to perform CRUD operations to manage a list of Books.
// Connects to the MongoDB session amd sets up the HTTP request routes using the gorilla/mux library.
// After defining the routes and the associated methods, the book api application runs and listens on port 8080.

package main

import (
	"github.com/gorilla/mux"
	"github.com/sbouaazzi/book/controllers"
	dao2 "github.com/sbouaazzi/book/dao"
	"log"
	"net/http"
)

// constants definitions
const (
	ApplicationPortNumberMsg = ":8080"
	ApplicationStartedMsg    = "Application Started"
	BookEndpoint             = "/book"
	BookIdParamEndpoint      = "/book/{id}"
	ContentType              = "Content-Type"
	DatabaseName             = "BookMongo"
	DatabaseServerUrl        = "mongodb://book_mongodb_1:27017"
	ListeningOnPort8080Msg   = "Listening on port 8080"
)

// BookDAO instance with MongoDB URL and 'BookMongo' Database
var dao = dao2.BookDAO{Server: DatabaseServerUrl, Database: DatabaseName}

// main function
//
// Defines HTTP request routes, connects to mongodb session and runs application server on port 8080
func main() {
	log.Println(ApplicationStartedMsg)
	// connect to MongoDB through the dao instance
	dao.Connect()
	r := mux.NewRouter()

	// GET
	r.HandleFunc(BookEndpoint, controllers.GetAllBooks).Methods(http.MethodGet)
	r.HandleFunc(BookIdParamEndpoint, controllers.GetBook).Methods(http.MethodGet)

	// POST
	r.HandleFunc(BookEndpoint, controllers.CreateBook).Methods(http.MethodPost)

	// PUT
	r.HandleFunc(BookIdParamEndpoint, controllers.UpdateBook).Methods(http.MethodPut)

	// DELETE
	r.HandleFunc(BookIdParamEndpoint, controllers.DeleteBook).Methods(http.MethodDelete)

	// run application and listen on port 8080
	log.Println(ListeningOnPort8080Msg)
	if err := http.ListenAndServe(ApplicationPortNumberMsg, r); err != nil {
		log.Fatal(err)
	}
}
