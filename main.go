package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// 1 - controller
func home(r http.ResponseWriter, q *http.Request) {
	if q.URL.Path != "/" {
		http.NotFound(r, q)
		return
	}
	r.Write([]byte("Hello World !"))
}

func createSnippet(r http.ResponseWriter, q *http.Request) {

	if q.Method != "POST" {
		r.Header().Set("Allow", "POST")
		r.Header()["X-XSS-Protection"] = []string{"1; mode=block"}
		r.Header()["Date"] = nil
		r.WriteHeader(405)
		r.Write([]byte("Method Not Allowed"))
		return
	}

	r.Write([]byte("create Snippet"))
}

func showSnippet(r http.ResponseWriter, q *http.Request) {
	id, err := strconv.Atoi(q.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(r, q)
		return
	}

	fmt.Fprintf(r, "Display snippet of ID %d", id)
}

func main() {
	// 2 - router
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Server Listinig on http://localhost:8080")
	// 3 - server
	err := http.ListenAndServe(":8080", mux)

	log.Fatal(err)
}
