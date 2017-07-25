package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", ListPost)
	http.ListenAndServe(":8000", mux)
}
