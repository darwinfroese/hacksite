package api

import (
	"net/http"
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

	rememberMe := r.URL.Query().Get("RememberMe")

	if !ok {
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, "BasicAuth failed")
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	session, err := auth.Login(*ctx.DB, username, password)

	if err != nil {
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, err.Error())
		if strings.Contains(err.Error(), auth.UnathorizedErrorMessage) {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	auth.SetCookie(w, auth.SessionCookieName, session.Token, rememberMe)
	w.WriteHeader(http.StatusOK)
}

func logoutHandler(ctx *Context, w http.ResponseWriter, r *http.Request) {
	sessionCookie, err := r.Cookie(auth.SessionCookieName)
	if err != nil {
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, err.Error())
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	err = (*ctx.DB).RemoveSession(sessionCookie.Value)

	if err != nil {
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
