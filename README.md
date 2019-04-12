# book API
This repository contains a CRUD API to manage a list of Books

## API Documentation
* __GET__       /book
    * retrieves all the books in the list of books
    * example request: localhost:8080/book
    * example **200** response:
        ```
        [
            {
                "id": "5cae21aad67494bbc5532bbc",
                "title": "Rat in the Hat",
                "author": "Mr. Suess",
                "publisher": "Random House Inc",
                "publishdate": "2019",
                "rating": 1,
                "status": "CheckedOut"
            },
            {
                "id": "5cae21aad67494bbc5532bbd",
                "title": "Bat in the Hat",
                "author": "Mrs. Suess",
                "publisher": "Random Planet Inc",
                "publishdate": "2015",
                "rating": 1,
                "status": "CheckedIn"
            } 
    ]
        ```
    * example **400** response:
        ```
            {
                error: "Invalid Book ID"
            }
        ```
* __GET__       /book/{id}
    * retrieves one book with the specified id parameter
    * example request: localhost:8080/book
    * example **200** response:
        ```
          {
              "id": "5cae21aad67494bbc5532bbc",
              "title": "Rat in the Hat",
              "author": "Mr. Suess",
              "publisher": "Random House Inc",
              "publishdate": "2019",
              "rating": 1,
              "status": "CheckedOut"
          }
        ```
    * example **400** response:
        ```
        {
            error: "Invalid Book ID"
        }
        ```
* __POST__      /book
    * creates one book with the data values in the message payload
    * example request: `localhost:8080/book`
        ```
            {
            	"title": "Nat in the hat",
            	"author": "Dr. Suess",
            	"publisher": "Random Planet",
            	"publishdate": "1960",
            	"rating": 2,
            	"status": "CheckedOut"
            }
        ```
    * example **200** response:
        ```
          {
              "id": "5cae21aad67494bbc5532bbc",
              "title": "Rat in the Hat",
              "author": "Mr. Suess",
              "publisher": "Random House Inc",
              "publishdate": "2019",
              "rating": 1,
              "status": "CheckedOut"
          }
        ```
    * example **400** response:
        ```
            {
                error: "Invalid Book ID"
            }
        ```
* __PUT__       /book/{id}
    * updates one book record with the specified id parameter in the book list
    * example request: `localhost:8080/book`
        ```
            {
               "title": "Nat in the hat",
               "author": "Dr. Suess",
               "publisher": "Random Planet",
               "publishdate": "1960",
               "rating": 2,
               "status": "CheckedOut"
            }
        ```
    * example **200** response:
       ```
          {
               "id": "5cae21aad67494bbc5532bbc",
               "title": "Rat in the Hat",
               "author": "Mr. Suess",
               "publisher": "Random House Inc",
               "publishdate": "2019",
               "rating": 1,
               "status": "CheckedOut"
          }
        ```
   * example **400** response:
        ```
           {
               error: "Invalid Book ID"
           }
        ```
* __DELETE__    /book/{id}
    * deletes one book record with the specified id parameter from the book list
    * example request: localhost:8080/book
    * example **200** response:
        ```
           {
             Result: Success
           }
        ```
    * example **400** response:
        ```
           {
             error: "Invalid Book ID"
           }
        ```

## Instructions
_Prerequisites: Be sure to have the Go environment and Docker installed on the local machine in order for the project to run properly_

In order to run this project from a docker container, please follow the following steps:

* Pull the project from the book github repo
* Open the project from any IDE
* From the terminal, navigate to the main /book directory
* In the terminal from that directory, run the command _docker-compose up --build_
* This command will build the MongoDB and Book AP Go source images and run the containers
* When the build is complete, and the containers are running, now access the API running on port 8080
* Use cURL or PostMan to test the API

