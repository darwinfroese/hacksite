package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/darwinfroese/hacksite/server/models"
)

func iterations(ctx context, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			addIteration(ctx, w, r)
			return
		case "OPTIONS":
			corsResponse(w, r)
			return
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
	}
}

func currentIteration(ctx context, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			switchIteration(ctx, w, r)
			return
		case "OPTIONS":
			corsResponse(w, r)
			return
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
	}
}

func addIteration(ctx context, w http.ResponseWriter, r *http.Request) {
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

func switchIteration(ctx context, w http.ResponseWriter, r *http.Request) {
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
