package api

import (
	"net/http"

	"github.com/darwinfroese/hacksite/server/database"
)

type context struct {
	db database.Database

	handler
}

func (ctx context) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx.handler(ctx, w, r)(w, r)
}
