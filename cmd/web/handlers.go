package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func home(r http.ResponseWriter, q *http.Request) {
	if q.URL.Path != "/" {
		http.NotFound(r, q)
		return
	}

	ts, err := template.ParseFiles("./ui/html/home.page.tmpl", "./ui/html/base.layout.tmpl", "./ui/html/footer.partial.tmpl")

	if err != nil {
		http.Error(r, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(r, nil)

	if err != nil {
		http.Error(r, "Internal Server Error", 500)
		return
	}
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
