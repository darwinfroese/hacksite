package api

import (
	"net/http"
)

// StatusWriter is a wrapper around ResponseWriter that stores the
// status code for usage
type statusWriter struct {
	http.ResponseWriter
	status  int
	written bool
}

// WriteHeader writes the header value for the router
func (w *statusWriter) WriteHeader(code int) {
	w.status = code
	if !w.written {
		w.ResponseWriter.WriteHeader(code)
		w.written = true
	}
}
