package api

import (
	"fmt"
	"net/http"
	"os"
)

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
func unsupportedMethodHandler(ctx Context, w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}