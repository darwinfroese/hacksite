package api

import (
	"net/http"
	"strings"
)

const (
	originAddress    = "http://localhost:8080"
	supportedHeaders = "Content-Type, Authorization"
)

// optionsHandler sets CORS headers and returns 200 -- used for OPTIONS requests
// TODO: This should be applied in an adapter
func optionsHandler(ctx apiContext, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", originAddress)
	w.Header().Set("Access-Control-Allow-Headers", supportedHeaders)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", methodsToString(ctx.supportedMethods))
	w.WriteHeader(http.StatusOK)
}

// unsupportedMethodHandler is a default handler that will send a 405 error
func unsupportedMethodHandler(ctx apiContext, w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}

func methodsToString(methods []string) string {
	return strings.Join(methods, ", ")
}

func callHandler(ctx apiContext, w http.ResponseWriter, r *http.Request, m map[string]handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if h, ok := m[r.Method]; ok {
			h(ctx, w, r)
		} else {
			unsupportedMethodHandler(ctx, w, r)
		}
	}
}
