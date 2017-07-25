package main

import (
	"fmt"
	"html"
	"net/http"
)

// ListPost write list of post to http
func ListPost(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, %q <br> ", html.EscapeString(req.URL.Path))
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(req.URL.Path))
}
