package api

import (
	"fmt"
	"net/http"
	"time"

	uuid "github.com/nu7hatch/gouuid"
)

const timeFormat = "2006-01-02 15:04:05"

type adapter func(*Context, http.HandlerFunc) http.HandlerFunc

var adapters = []adapter{
	setCorsHeaders,
	logRequestDuration,
}

// Apply will call all middleware functions on the incoming handler request
func Apply(ctx *Context, h http.HandlerFunc) http.HandlerFunc {
	for _, a := range adapters {
		h = a(ctx, h)
	}

	return h
}

func logRequestDuration(ctx *Context, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx.RequestID = getUUID()

		start := time.Now()

		(*ctx.Logger).InfoWithRequest(r, ctx.RequestID, "Request started")
		defer func() {
			dur := fmt.Sprintf("Request duration: %s", time.Since(start).String())
			(*ctx.Logger).InfoWithRequest(r, ctx.RequestID, dur)
		}()

		// This middleware is hit first
		sw := statusWriter{ResponseWriter: w, status: 200}

		h(&sw, r)
	}
}

func setCorsHeaders(ctx *Context, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", originAddress)
		w.Header().Set("Access-Control-Allow-Headers", supportedHeaders)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, PATCH, DELETE")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		h(w, r)
	}
}

func getUUID() string {
	uid := ""
	id, err := uuid.NewV4()

	if err == nil {
		uid = id.String()
	}

	return uid
}
