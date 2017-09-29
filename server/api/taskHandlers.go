package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/darwinfroese/hacksite/server/models"
)

func tasks(ctx context, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "PUT":
			updateTask(ctx, w, r)
			return
		case "PATCH":
			updateTask(ctx, w, r)
			return
		case "DELETE":
			removeTask(ctx, w, r)
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

// Handlers for specific methods on /tasks
func updateTask(ctx context, w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	project, err := ctx.db.UpdateTask(task)
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

func removeTask(ctx context, w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	project, err := ctx.db.RemoveTask(task)
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
