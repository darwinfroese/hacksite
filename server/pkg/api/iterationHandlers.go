package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

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
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	var iteration models.Iteration
	err := json.NewDecoder(r.Body).Decode(&iteration)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	project, err := iterations.CreateIteration(*ctx.DB, iteration)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(project)
}

func switchIteration(ctx *Context, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	var iteration models.Iteration
	err := json.NewDecoder(r.Body).Decode(&iteration)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	project, err := iterations.SwapCurrentIteration(*ctx.DB, iteration)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(project)
}
