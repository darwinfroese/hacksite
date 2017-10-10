package api

import (
	"net/http"
	"strings"
)

const (
	originAddress    = "http://localhost:8080"
	supportedHeaders = "Content-Type, Authorization"
)

type handler func(Context, http.ResponseWriter, *http.Request)

func methodsToString(methods []string) string {
	return strings.Join(methods, ", ")
}

func callHandler(ctx Context, w http.ResponseWriter, r *http.Request, m map[string]handler) {
	if h, ok := m[r.Method]; ok {
		h(ctx, w, r)
	} else {
		unsupportedMethodHandler(ctx, w, r)
	}
}
