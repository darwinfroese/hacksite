package api

import (
	"net/http"
	"time"

	"github.com/darwinfroese/hacksite/server/pkg/auth"
)

var sessionHandlersMap = map[string]handler{
	"GET": sessionHandler,
}

func (ctx *Context) sessionRoute(w http.ResponseWriter, r *http.Request) {
	callHandler(ctx, w, r, sessionHandlersMap)
}

func sessionHandler(ctx *Context, w http.ResponseWriter, r *http.Request) {
	session, err := auth.GetCurrentSession(*ctx.DB, *ctx.Logger, r)
	if err != nil {
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, err.Error())
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	if time.Now().After(session.Expiration) {
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, "Session was expired")
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	session.Expiration = time.Now().Add(time.Second * auth.SessionMaxAge)
	err = (*ctx.DB).StoreSession(session)
	if err != nil {
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	auth.SetCookie(w, auth.SessionCookieName, session.Token)

	w.WriteHeader(http.StatusOK)
}
