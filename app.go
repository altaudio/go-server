package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func logging() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(responseWriter http.ResponseWriter, request *http.Request) {
			start := time.Now()
			defer func() { log.Println(request.URL.Path, time.Since(start)) }()

			next(responseWriter, request)
		}
	}
}

func chainMiddlewares(handler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}

	return handler
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func main() {
	port := ":8080"
	router := mux.NewRouter()

	router.HandleFunc("/", chainMiddlewares(hello, logging()))

	fmt.Printf("Server listening on port%s", port)
	http.ListenAndServe(port, router)
}
