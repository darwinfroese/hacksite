package api

import (
	"fmt"
	"net/http"

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

func login(ctx apiContext, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return callHandler(ctx, w, r, loginHandlersMap)
}

func logout(ctx apiContext, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return callHandler(ctx, w, r, logoutHandlersMap)
}

func loginHandler(ctx apiContext, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	username, password, ok := r.BasicAuth()

	if !ok {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	account, err := ctx.db.GetAccount(username)
	if err != nil {
		if err.Error() == "no matching account found" {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		} else {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	password, err = auth.GetSaltedPassword(password, account.Salt)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if password != account.Password {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	session := auth.CreateSession(account.ID)
	err = ctx.db.StoreSession(session)

	if err != nil {
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
		fmt.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	err = ctx.db.RemoveSession(sessionCookie.Value)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
