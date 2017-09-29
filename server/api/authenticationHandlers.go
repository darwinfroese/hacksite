package api

import (
	"fmt"
	"net/http"

	"github.com/darwinfroese/hacksite/server/utilities"
)

func login(ctx context, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			loginHandler(ctx, w, r)
			return
		case "OPTIONS":
			corsResponse(w, r)
			return
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
	}
}

func logout(ctx context, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			logoutHandler(ctx, w, r)
			return
		case "OPTIONS":
			corsResponse(w, r)
			return
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
	}
}

// TODO: Extend session length everytime the user interacts with the API
func loginHandler(ctx context, w http.ResponseWriter, r *http.Request) {
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

	password, err = utilities.GetSaltedPassword(password, account.Salt)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if password != account.Password {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	session := utilities.CreateSession(account.ID)
	err = ctx.db.StoreSession(session)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	utilities.SetCookie(w, utilities.SessionCookieName, session.Token)
	w.WriteHeader(http.StatusOK)
}

func logoutHandler(ctx context, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	sessionCookie, err := r.Cookie(utilities.SessionCookieName)
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
