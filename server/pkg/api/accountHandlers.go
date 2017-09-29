package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/darwinfroese/hacksite/server/models"
	"github.com/darwinfroese/hacksite/server/pkg/auth"
)

var accountHandlerMap = map[string]handler{
	"POST":    createAccount,
	"OPTIONS": optionsHandler,
}

func accounts(ctx apiContext, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return callHandler(ctx, w, r, accountHandlerMap)
}

func createAccount(ctx apiContext, w http.ResponseWriter, r *http.Request) {
	var account models.Account
	err := json.NewDecoder(r.Body).Decode(&account)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	salt, password, err := auth.SaltPassword(account.Password)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	account.Password = password
	account.Salt = salt

	_, err = ctx.db.CreateAccount(account)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// TODO: Make a helper function for this
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.WriteHeader(http.StatusCreated)
}
