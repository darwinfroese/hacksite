package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/darwinfroese/hacksite/server/models"
)

type message struct {
	Message string
}

type handler func(context, http.ResponseWriter, *http.Request) http.HandlerFunc

func projects(ctx context, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Method:", r.Method)
		switch r.Method {
		case "GET":
			getAllProjects(ctx, w)
			return
		case "POST":
			createProject(ctx, w, r)
			return
		case "DELETE":
			deleteProject(ctx, w, r)
			return
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
	}
}

func getProject(ctx context, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		args := r.URL.Query()
		str := args.Get("id")

		if str == "" {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(str)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		msg := fmt.Sprintf("Get Project %d", id)

		ctx.db.GetProject(id)
		json.NewEncoder(w).Encode(message{Message: msg})
	}
}

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
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
	}
}

// Handlers for specific methods on /projects
func getAllProjects(ctx context, w http.ResponseWriter) {
	ctx.db.GetProjects()
	json.NewEncoder(w).Encode(message{Message: "Get All Projects"})
}

func createProject(ctx context, w http.ResponseWriter, r *http.Request) {
	var project models.Project
	err := json.NewDecoder(r.Body).Decode(&project)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	ctx.db.AddProject(project)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(message{Message: "Created a new project"})
}

func deleteProject(ctx context, w http.ResponseWriter, r *http.Request) {
	var project models.Project
	err := json.NewDecoder(r.Body).Decode(&project)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	msg := fmt.Sprintf("Project %d deleted", project.ID)

	ctx.db.RemoveProject(project.ID)
	json.NewEncoder(w).Encode(message{Message: msg})
}

// Handlers for specific methods on /tasks
func updateTask(ctx context, w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	ctx.db.UpdateTask(task)
	json.NewEncoder(w).Encode(message{Message: "Task Updated"})
}

func removeTask(ctx context, w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	ctx.db.RemoveTask(task)
	json.NewEncoder(w).Encode(message{Message: "Task Deleted"})
}
