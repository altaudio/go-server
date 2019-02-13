package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	port := ":8080"
	router := mux.NewRouter()

	template, error := template.ParseFiles("template.html")

	if error != nil {
		log.Fatal(error)
	}

	router.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
		template.Execute(responseWriter, struct{ LinkClicked bool }{false})
	})

	router.HandleFunc("/click", func(responseWriter http.ResponseWriter, request *http.Request) {
		template.Execute(responseWriter, struct{ LinkClicked bool }{true})
	})

	fmt.Printf("Server listening on port%s", port)
	http.ListenAndServe(port, router)
}
