package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// 2 - router
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	flag.Parse()
	log.Println("Server Listinig on http://localhost:8080")
	// 3 - server
	err := http.ListenAndServe(":8080", mux)

	log.Fatal(err)
}
