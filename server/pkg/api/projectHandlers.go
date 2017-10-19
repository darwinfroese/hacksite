package api

import (
	"encoding/json"
	"net/http"

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
	id := args.Get("id")

	if id == "" {
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, "Bad request received")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	project, err := (*ctx.DB).GetProject(id)
	if err != nil {
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(project)
}

// Handlers for specific methods on /projects
func getAllProjects(ctx *Context, w http.ResponseWriter, r *http.Request) {
	session, err := auth.GetCurrentSession(*ctx.DB, *ctx.Logger, r)
	if err != nil {
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, "Unauthorized access")
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	projects, err := projects.GetUserProjects(*ctx.DB, *ctx.Logger, session.Username)
	if err != nil {
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(projects)
}

func createProject(ctx *Context, w http.ResponseWriter, r *http.Request) {
	var project models.Project
	err := json.NewDecoder(r.Body).Decode(&project)

	if err != nil {
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = project.Validate()
	if err != nil {
		// TODO: Make this return an error (validate)
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, "Project was invalid")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	session, err := auth.GetCurrentSession(*ctx.DB, *ctx.Logger, r)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	err = projects.CreateProject(*ctx.DB, *ctx.Logger, &project, session.Username)
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
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, "Bad project request")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = projects.UpdateProject(*ctx.DB, &project)
	if err != nil {
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, err.Error())
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
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	session, err := auth.GetCurrentSession(*ctx.DB, *ctx.Logger, r)
	if err != nil {
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, err.Error())
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	err = projects.DeleteProject(*ctx.DB, *ctx.Logger, session.Username, project.ID)
	if err != nil {
		(*ctx.Logger).ErrorWithRequest(r, ctx.RequestID, err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
