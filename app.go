package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	port := ":8080"
	router := mux.NewRouter()

	// Example of getting the URL segments
	router.HandleFunc("/books/{title}/page/{page}", func(responseWriter http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(responseWriter, "You've requested the book %s at page %s", title, page)
	})

	// Example of restricting request handlers to HTTP methods
	router.HandleFunc("/books", func(responseWriter http.ResponseWriter, request *http.Request) {
		fmt.Fprint(responseWriter, "POST /books Response")
	}).Methods("POST")

	// Subrouter Example
	magazineRouter := router.PathPrefix("/magazines").Subrouter()
	magazineRouter.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
		fmt.Fprint(responseWriter, "GET /magazines Response")
	})
	magazineRouter.HandleFunc("/{title}", func(responseWriter http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		title := vars["title"]

		fmt.Fprintf(responseWriter, "GET /magazines title: %s", title)
	})

	fmt.Printf("Server listening on port%s", port)
	http.ListenAndServe(port, router)
}
