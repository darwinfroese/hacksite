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

	apiErr := accounts.CreateAccount(*ctx.DB, *ctx.Logger, &account)

	if apiErr != nil {
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, apiErr.Error())
		http.Error(w, apiErr.Message, apiErr.Code)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(account)
}
