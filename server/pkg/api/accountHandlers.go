package api

import (
	"encoding/json"
	"net/http"

	"github.com/darwinfroese/hacksite/server/models"
	"github.com/darwinfroese/hacksite/server/pkg/accounts"
	"github.com/darwinfroese/hacksite/server/pkg/auth"
)

var accountsHandlerMap = map[string]handler{
	"POST": createAccount,
}

var accountHandlerMap = map[string]handler{
	"GET": getAccount,
	"POST": updateAccount,
}

func (ctx *Context) accountsRoute(w http.ResponseWriter, r *http.Request) {
	callHandler(ctx, w, r, accountsHandlerMap)
}

func (ctx *Context) accountRoute(w http.ResponseWriter, r *http.Request) {
	callHandler(ctx, w, r, accountHandlerMap)
}

func getAccount(ctx *Context, w http.ResponseWriter, r *http.Request) {
	var account models.Account

	session, err := auth.GetCurrentSession(*ctx.DB, *ctx.Logger, r)
	if err != nil {
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, err.Error())
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	account, err = (*ctx.DB).GetAccountByUsername(session.Username)
	if err != nil {
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(account)
}

func updateAccount(ctx *Context, w http.ResponseWriter, r *http.Request) {
	var updateData models.Account
	var account models.Account

	err := json.NewDecoder(r.Body).Decode(&updateData)

	if err != nil {
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	session, err := auth.GetCurrentSession(*ctx.DB, *ctx.Logger, r)
	if err != nil {
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, err.Error())
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	account, err = (*ctx.DB).GetAccountByUsername(session.Username)
	if err != nil {
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	account.Password = updateData.Password

	json.NewEncoder(w).Encode(account)
	//(*ctx.DB).UpdateAccount(account)
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
