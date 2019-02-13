package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	port := ":8080"
	router := mux.NewRouter()

	data := TodoPageData{
		PageTitle: "Todo List",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}

	template, error := template.ParseFiles("template.html")

	if error != nil {
		log.Fatal(error)
	}

	router.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
		template.Execute(responseWriter, data)
	})

	fmt.Printf("Server listening on port%s", port)
	http.ListenAndServe(port, router)
}
