package api

import (
	"net/http"

	"github.com/darwinfroese/hacksite/server/database"
)

type apiHandler func(apiContext, http.ResponseWriter, *http.Request) http.HandlerFunc
type handler func(apiContext, http.ResponseWriter, *http.Request)

type apiContext struct {
	db               database.Database
	supportedMethods []string

	apiHandler
}

func (ctx apiContext) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx.apiHandler(ctx, w, r)(w, r)
}
