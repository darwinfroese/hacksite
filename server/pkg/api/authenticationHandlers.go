package api

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/darwinfroese/hacksite/server/pkg/auth"
)

var loginHandlersMap = map[string]handler{
	"GET": loginHandler,
}

var logoutHandlersMap = map[string]handler{
	"GET": logoutHandler,
}

func (ctx *Context) loginRoute(w http.ResponseWriter, r *http.Request) {
	callHandler(ctx, w, r, loginHandlersMap)
}

func (ctx *Context) logoutRoute(w http.ResponseWriter, r *http.Request) {
	callHandler(ctx, w, r, logoutHandlersMap)
}

func loginHandler(ctx *Context, w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()

	if !ok {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	session, err := auth.Login(*ctx.DB, username, password)

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

func logoutHandler(ctx *Context, w http.ResponseWriter, r *http.Request) {
	sessionCookie, err := r.Cookie(auth.SessionCookieName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	err = (*ctx.DB).RemoveSession(sessionCookie.Value)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
