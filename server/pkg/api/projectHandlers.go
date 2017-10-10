package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/darwinfroese/hacksite/server/models"
	"github.com/darwinfroese/hacksite/server/pkg/auth"
	"github.com/darwinfroese/hacksite/server/pkg/projects"
)

// TODO: Access-Control-Allow-Origin needs to restrict to production web port
// TODO: A lot of these handlers are very similar
// TODO: Make sure correct status codes are being returned
// TODO: Set CORS earlier since it "errors" if http.Error happens
// TODO: Web-client handles auth right now, should be implemented into middleware

var projectHandlersMap = map[string]handler{
	"GET": getProject,
}

var projectsHandlersMap = map[string]handler{
	"GET":    getAllProjects,
	"POST":   createProject,
	"PUT":    updateProject,
	"PATCH":  updateProject,
	"DELETE": deleteProject,
}

func (ctx *Context) projectRoute(w http.ResponseWriter, r *http.Request) {
	callHandler(ctx, w, r, projectHandlersMap)
}

func (ctx *Context) projectsRoute(w http.ResponseWriter, r *http.Request) {
	callHandler(ctx, w, r, projectsHandlersMap)
}

func getProject(ctx *Context, w http.ResponseWriter, r *http.Request) {
	args := r.URL.Query()
	str := args.Get("id")

	if str == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(str)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	project, err := (*ctx.DB).GetProject(id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(project)
}

// Handlers for specific methods on /projects
func getAllProjects(ctx *Context, w http.ResponseWriter, r *http.Request) {
	session, err := auth.GetCurrentSession(*ctx.DB, r)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	projects, err := projects.GetUserProjects(*ctx.DB, session.UserID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(projects)
}

func createProject(ctx *Context, w http.ResponseWriter, r *http.Request) {
	var project models.Project
	err := json.NewDecoder(r.Body).Decode(&project)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if !projects.ValidateProject(project) {
		// TODO: Make this return an error (validate)
		fmt.Fprintf(os.Stderr, "Error: %s\n", "Project was invalid!")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	session, err := auth.GetCurrentSession(*ctx.DB, r)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	err = projects.CreateProject(*ctx.DB, &project, session)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(project)
}

func updateProject(ctx *Context, w http.ResponseWriter, r *http.Request) {
	// TODO: this can be refactored
	var project models.Project
	err := json.NewDecoder(r.Body).Decode(&project)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = projects.UpdateProject(*ctx.DB, &project)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(project)
}

func deleteProject(ctx *Context, w http.ResponseWriter, r *http.Request) {
	var project models.Project
	err := json.NewDecoder(r.Body).Decode(&project)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	session, err := auth.GetCurrentSession(*ctx.DB, r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	err = projects.DeleteProject(*ctx.DB, session.UserID, project.ID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
