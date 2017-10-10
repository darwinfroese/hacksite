package api

import (
	"fmt"
	"net/http"
	"os"
	"time"

	uuid "github.com/nu7hatch/gouuid"
)

const timeFormat = "2006-01-02 15:04:05"

type adapter func(Context, http.HandlerFunc) http.HandlerFunc

var adapters = []adapter{
	logRequestDuration,
	setCorsHeaders,
}

// Apply will call all middleware functions on the incoming handler request
func Apply(ctx Context, h http.HandlerFunc) http.HandlerFunc {
	for _, a := range adapters {
		h = a(ctx, h)
	}

	return h
}

func logRequestDuration(ctx Context, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx.RequestID = getUUID()

		start := time.Now()
		fmt.Fprintf(os.Stdout, "[ INFO ] %s :: Request[ID: %s] to %s started.\n",
			start.Format(timeFormat), ctx.RequestID, r.URL.Path)

		defer func() {
			fmt.Fprintf(os.Stdout, "[ INFO ] %s :: Request[ID: %s] to %s finished (duration: %s).\n",
				time.Now().Format(timeFormat), ctx.RequestID, r.URL.Path, time.Since(start).String())
		}()

		h(w, r)
	}
}

func setCorsHeaders(ctx Context, h http.HandlerFunc) http.HandlerFunc {
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
