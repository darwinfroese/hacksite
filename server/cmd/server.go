package main

import (
	"fmt"
	"net/http"
)

const (
	version = "1.0.0"
)

func main() {
	m := http.NewServeMux()

	m.Handle("/", http.FileServer(http.Dir("./webdist")))
	m.HandleFunc("/api/v1/version", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API Version: %s", version)
	})

	http.ListenAndServe(":8800", m)
}
