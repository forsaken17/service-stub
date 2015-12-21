package main

import (
	"io"
	"net/http"
)

func dummy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8") // normal header
	io.WriteString(w, `{"dummy":"empty"}`)
}

func main() {
	http.HandleFunc("/", dummy)
	http.ListenAndServe(":9200", nil)
}
