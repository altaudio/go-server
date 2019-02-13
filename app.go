package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := ":8080"
	staticDirectory := "static/"
	staticURL := "/static/"

	fileServer := http.FileServer(http.Dir(staticDirectory))

	http.Handle(staticURL, http.StripPrefix(staticURL, fileServer))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website: %s", r.URL.Query().Get("name"))
	})

	fmt.Printf("Listening on port%s\n", port)
	http.ListenAndServe(port, nil)
}
