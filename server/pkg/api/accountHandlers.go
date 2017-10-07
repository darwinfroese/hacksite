package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"


	"github.com/darwinfroese/hacksite/server/models"
	"github.com/darwinfroese/hacksite/server/pkg/accounts"
)


var accountHandlerMap = map[string]handler{
	"POST":    createAccount,
	"OPTIONS": optionsHandler,
}

func accountsRoute(ctx apiContext, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return callHandler(ctx, w, r, accountHandlerMap)
}

func createAccount(ctx apiContext, w http.ResponseWriter, r *http.Request) {
	// TODO: Make a helper function for this
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	var account models.Account
	err := json.NewDecoder(r.Body).Decode(&account)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = accounts.CreateAccount(ctx.db, &account)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		if err.Error() == models.EmailTakenErrorMessage || err.Error() == models.UsernameTakenErrorMessage {
			res := models.ResponseObject{http.StatusConflict, err.Error(), ""}
			json.NewEncoder(w).Encode(res)
			return
		} else {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(account)
}
