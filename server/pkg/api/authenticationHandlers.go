package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/darwinfroese/hacksite/server/pkg/auth"
)

type loginRequest struct {
	RememberMe bool
}

var loginHandlersMap = map[string]handler{
	"GET":  loginHandler,
	"POST": postLoginHandler,
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
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, "BasicAuth failed")
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	rememberMe, err := strconv.ParseBool(r.URL.Query().Get("RememberMe"))
	if err != nil {
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, "Remember Me failed")
		rememberMe = false
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

func postLoginHandler(ctx *Context, w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()

	if !ok {
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, "BasicAuth failed")
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	var content loginRequest
	err := json.NewDecoder(r.Body).Decode(&content)

	if err != nil {
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, "Remember Me failed")
		content.RememberMe = false
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

	auth.SetCookie(w, auth.SessionCookieName, session.Token, content.RememberMe)
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
