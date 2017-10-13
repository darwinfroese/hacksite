package api

import (
	"encoding/json"
	"net/http"

	"github.com/darwinfroese/hacksite/server/models"
	"github.com/darwinfroese/hacksite/server/pkg/evolutions"
)

var iterHandlersMap = map[string]handler{
	"POST": addEvolution,
}

var currIterHandlersMap = map[string]handler{
	"POST": switchEvolution,
}

func (ctx *Context) evolutionsRoute(w http.ResponseWriter, r *http.Request) {
	callHandler(ctx, w, r, iterHandlersMap)
}

func (ctx *Context) currentEvolutionRoute(w http.ResponseWriter, r *http.Request) {
	callHandler(ctx, w, r, currIterHandlersMap)
}

func addEvolution(ctx *Context, w http.ResponseWriter, r *http.Request) {
	var evolution models.Evolution
	err := json.NewDecoder(r.Body).Decode(&evolution)

	if err != nil {
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	project, err := evolutions.CreateEvolution(*ctx.DB, *ctx.Logger, evolution)

	if err != nil {
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(project)
}

func switchEvolution(ctx *Context, w http.ResponseWriter, r *http.Request) {
	var evolution models.Evolution
	err := json.NewDecoder(r.Body).Decode(&evolution)

	if err != nil {
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	project, err := evolutions.SwapCurrentEvolution(*ctx.DB, *ctx.Logger, evolution)
	if err != nil {
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(project)
}
