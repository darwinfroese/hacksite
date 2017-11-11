package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/darwinfroese/hacksite/server/models"
	"github.com/darwinfroese/hacksite/server/pkg/accounts"
	"github.com/darwinfroese/hacksite/server/pkg/auth"
)

type accountDto struct {
	Username, Name, Email string
}

var accountHandlerMap = map[string]handler{
	"GET":   getAccountInfo,
	"POST":  createAccount,
	"PUT":   updateAccount,
	"PATCH": updateAccount,
}

func (ctx *Context) accountsRoute(w http.ResponseWriter, r *http.Request) {
	callHandler(ctx, w, r, accountHandlerMap)
}

func getAccountInfo(ctx *Context, w http.ResponseWriter, r *http.Request) {
	sesh, err := auth.GetCurrentSession(*ctx.DB, *ctx.Logger, r)
	if err != nil {
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, err.Error())
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	acc, err := auth.GetCurrentAccountFromSession(*ctx.DB, *ctx.Logger, sesh)
	if err != nil {
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, err.Error())
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accountDto{
		Username: acc.Username,
		Email:    acc.Email,
		Name:     acc.Name,
	})
}

func createAccount(ctx *Context, w http.ResponseWriter, r *http.Request) {
	var account models.Account
	err := json.NewDecoder(r.Body).Decode(&account)

	if err != nil {
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	account.Username = strings.ToLower(account.Username)
	apiErr := accounts.CreateAccount(*ctx.DB, *ctx.Logger, &account)

	if apiErr != nil {
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, apiErr.FullError())
		http.Error(w, apiErr.Error(), apiErr.Code)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(account)
}

func updateAccount(ctx *Context, w http.ResponseWriter, r *http.Request) {
	var account models.Account
	err := json.NewDecoder(r.Body).Decode(&account)

	if err != nil {
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	account.Username = strings.ToLower(account.Username)
	sesh, err := auth.GetCurrentSession(*ctx.DB, *ctx.Logger, r)
	if err != nil {
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, err.Error())
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	acc, err := auth.GetCurrentAccountFromSession(*ctx.DB, *ctx.Logger, sesh)
	if err != nil {
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, err.Error())
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	apiErr := accounts.UpdateAccount(*ctx.DB, acc.Username, acc.Email, account)
	if apiErr != nil {
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, apiErr.FullError())
		http.Error(w, apiErr.Error(), apiErr.Code)
		return
	}

	sesh.Username = account.Username
	err = (*ctx.DB).StoreSession(sesh)
	if err != nil {
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(account)
}
