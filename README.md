# Go Server Test
## References

- [A Tour of Go](https://tour.golang.org)
- [Go Web Examples](https://gowebexamples.com)
- [Gorilla Mux Docs](https://github.com/gorilla/mux)
- [Realize](https://github.com/oxequa/realize)
- [Build a web application with Go](https://astaxie.gitbooks.io/build-web-application-with-golang/en/01.0.html)
- [Go Docs](https://golang.org/doc/)

## Setup
- [Install Go](https://golang.org/doc/install)
- To get global go packages working add `export PATH=$HOME/go/bin:$PATH` to `~/.bash_profile`
- `go get github.com/oxequa/realize`

## Development
- Run `realize start`

## Examples
### [Basic Server with Static Folder](https://gowebexamples.com/http-server/)
```
// app.go
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
```

### [Gorilla Mux Router Example](https://gowebexamples.com/routes-using-gorilla-mux/)
```
// app.go
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
```
### [Hacky Interactive Templates Example](https://gowebexamples.com/forms/)
```
// app.go

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

```
```
// template.html

{{if .LinkClicked}}
	<h1>Link Clicked</h1>
{{else}}
    <a href="/click">Click Me</a>
{{end}}
```