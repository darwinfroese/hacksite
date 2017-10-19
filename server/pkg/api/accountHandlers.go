package api

import (
	"encoding/json"
	"net/http"

	"github.com/darwinfroese/hacksite/server/models"
	"github.com/darwinfroese/hacksite/server/pkg/accounts"
)

var accountHandlerMap = map[string]handler{
	"POST": createAccount,
}

func (ctx *Context) accountsRoute(w http.ResponseWriter, r *http.Request) {
	callHandler(ctx, w, r, accountHandlerMap)
}

func createAccount(ctx *Context, w http.ResponseWriter, r *http.Request) {
	var account models.Account
	err := json.NewDecoder(r.Body).Decode(&account)

	if err != nil {
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = accounts.CreateAccount(*ctx.DB, *ctx.Logger, &account)

	if err != nil {
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, err.Error())
		if err.Error() == models.EmailTakenErrorMessage || err.Error() == models.UsernameTakenErrorMessage {
			res := models.ResponseObject{
				StatusCode:   http.StatusConflict,
				ErrorMessage: err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(account)
}
