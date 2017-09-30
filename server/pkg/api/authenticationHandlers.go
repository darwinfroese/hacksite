package api

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/darwinfroese/hacksite/server/pkg/auth"
)

var loginHandlersMap = map[string]handler{
	"GET":     loginHandler,
	"OPTIONS": optionsHandler,
}

var logoutHandlersMap = map[string]handler{
	"GET":     logoutHandler,
	"OPTIONS": optionsHandler,
}

func loginRoute(ctx apiContext, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return callHandler(ctx, w, r, loginHandlersMap)
}

func logoutRoute(ctx apiContext, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return callHandler(ctx, w, r, logoutHandlersMap)
}

func loginHandler(ctx apiContext, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	username, password, ok := r.BasicAuth()

	if !ok {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	session, err := auth.Login(ctx.db, username, password)

	if err != nil {
		if strings.Contains(err.Error(), auth.UnathorizedErrorMessage) {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	auth.SetCookie(w, auth.SessionCookieName, session.Token)
	w.WriteHeader(http.StatusOK)
}

func logoutHandler(ctx apiContext, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	sessionCookie, err := r.Cookie(auth.SessionCookieName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	err = ctx.db.RemoveSession(sessionCookie.Value)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
