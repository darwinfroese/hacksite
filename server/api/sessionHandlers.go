package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/darwinfroese/hacksite/server/utilities"
)

func session(ctx context, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			sessionHandler(ctx, w, r)
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

func sessionHandler(ctx context, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	sessionCookie, err := r.Cookie(utilities.SessionCookieName)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	session, err := ctx.db.GetSession(sessionCookie.Value)
	if time.Now().After(session.Expiration) {
		fmt.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	session.Expiration = time.Now().Add(time.Second * utilities.SessionMaxAge)
	err = ctx.db.StoreSession(session)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	utilities.SetCookie(w, utilities.SessionCookieName, session.Token)

	w.WriteHeader(http.StatusOK)
}
