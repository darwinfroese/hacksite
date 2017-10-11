package api

import (
	"encoding/json"
	"net/http"

	"github.com/darwinfroese/hacksite/server/models"
	"github.com/darwinfroese/hacksite/server/pkg/iterations"
)

var iterHandlersMap = map[string]handler{
	"POST": addIteration,
}

var currIterHandlersMap = map[string]handler{
	"POST": switchIteration,
}

func (ctx *Context) iterationsRoute(w http.ResponseWriter, r *http.Request) {
	callHandler(ctx, w, r, iterHandlersMap)
}

func (ctx *Context) currentIterationRoute(w http.ResponseWriter, r *http.Request) {
	callHandler(ctx, w, r, currIterHandlersMap)
}

func addIteration(ctx *Context, w http.ResponseWriter, r *http.Request) {
	var iteration models.Iteration
	err := json.NewDecoder(r.Body).Decode(&iteration)

	if err != nil {
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	project, err := iterations.CreateIteration(*ctx.DB, *ctx.Logger, iteration)

	if err != nil {
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(project)
}

func switchIteration(ctx *Context, w http.ResponseWriter, r *http.Request) {
	var iteration models.Iteration
	err := json.NewDecoder(r.Body).Decode(&iteration)

	if err != nil {
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	project, err := iterations.SwapCurrentIteration(*ctx.DB, *ctx.Logger, iteration)
	if err != nil {
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(project)
}
