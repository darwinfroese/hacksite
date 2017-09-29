package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/darwinfroese/hacksite/server/models"
)

var iterHandlersMap = map[string]handler{
	"POST":    addIteration,
	"OPTIONS": optionsHandler,
}

var currIterHandlersMap = map[string]handler{
	"POST":    switchIteration,
	"OPTOINS": optionsHandler,
}

func iterations(ctx apiContext, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return callHandler(ctx, w, r, iterHandlersMap)
}

func currentIteration(ctx apiContext, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return callHandler(ctx, w, r, iterHandlersMap)
}

func addIteration(ctx apiContext, w http.ResponseWriter, r *http.Request) {
	var iteration models.Iteration
	err := json.NewDecoder(r.Body).Decode(&iteration)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	project, err := ctx.db.AddIteration(iteration)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	json.NewEncoder(w).Encode(project)
}

func switchIteration(ctx apiContext, w http.ResponseWriter, r *http.Request) {
	var iteration models.Iteration
	err := json.NewDecoder(r.Body).Decode(&iteration)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	project, err := ctx.db.SwapCurrentIteration(iteration)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	json.NewEncoder(w).Encode(project)
}
