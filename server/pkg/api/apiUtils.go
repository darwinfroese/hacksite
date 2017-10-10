package api

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

const (
	originAddress    = "http://localhost:8080"
	supportedHeaders = "Content-Type, Authorization"
)

type handler func(*Context, http.ResponseWriter, *http.Request)

// optionsHandler sets CORS headers and returns 200 -- used for OPTIONS requests
// TODO: This should be applied in an adapter
func optionsHandler(ctx *Context, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", originAddress)
	w.Header().Set("Access-Control-Allow-Headers", supportedHeaders)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, PATCH, DELETE")
	w.WriteHeader(http.StatusOK)
}

// RedirectToHTTPS will take the request coming in on *:80 and forward it to *:443
func RedirectToHTTPS(w http.ResponseWriter, r *http.Request) {
	target := "https://" + r.Host + r.URL.Path

	if len(r.URL.RawQuery) > 0 {
		target += "?" + r.URL.RawQuery
	}

	fmt.Fprintf(os.Stdout, "Redirecting to: %s\n", target)
	http.Redirect(w, r, target, http.StatusTemporaryRedirect)
}

// healthCheckHandler responds with a simple "OK" to say the server is in a good state
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

// unsupportedMethodHandler is a default handler that will send a 405 error
func unsupportedMethodHandler(ctx *Context, w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}

func methodsToString(methods []string) string {
	return strings.Join(methods, ", ")
}

func callHandler(ctx *Context, w http.ResponseWriter, r *http.Request, m map[string]handler) {
	if h, ok := m[r.Method]; ok {
		h(ctx, w, r)
	} else {
		unsupportedMethodHandler(ctx, w, r)
	}
}
