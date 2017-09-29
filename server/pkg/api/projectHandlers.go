package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/darwinfroese/hacksite/server/models"
	"github.com/darwinfroese/hacksite/server/pkg/auth"
	"github.com/darwinfroese/hacksite/server/pkg/projects"
)

// TODO: Access-Control-Allow-Origin needs to restrict to production web port
// TODO: move switch/case into a function that can be shared
// TODO: A lot of these handlers are very similar
// TODO: Make sure correct status codes are being returned
// TODO: Set CORS earlier since it "errors" if http.Error happens

var projectHandlersMap = map[string]handler{
	"GET":     getProject,
	"OPTIONS": optionsHandler,
}

var projectsHandlersMap = map[string]handler{
	"GET":     getAllProjects,
	"POST":    createProject,
	"PUT":     updateProject,
	"PATCH":   updateProject,
	"DELETE":  deleteProject,
	"OPTIONS": optionsHandler,
}

func project(ctx apiContext, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return callHandler(ctx, w, r, projectHandlersMap)
}

func projectsRoute(ctx apiContext, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return callHandler(ctx, w, r, projectsHandlersMap)
}

func getProject(ctx apiContext, w http.ResponseWriter, r *http.Request) {
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

	project, err := ctx.db.GetProject(id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(project)
}

// Handlers for specific methods on /projects
func getAllProjects(ctx apiContext, w http.ResponseWriter, r *http.Request) {
	sessionToken, err := r.Cookie(auth.SessionCookieName)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	session, err := ctx.db.GetSession(sessionToken.Value)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	if time.Now().After(session.Expiration) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	projects, err := ctx.db.GetProjects(session.UserID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}

func createProject(ctx apiContext, w http.ResponseWriter, r *http.Request) {
	var project models.Project
	err := json.NewDecoder(r.Body).Decode(&project)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if !projects.ValidateProject(project) {
		// TODO: Make this return an error
		fmt.Fprintf(os.Stderr, "Error: %s\n", "Project was invalid!")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	project, err = ctx.db.AddProject(project)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// TODO: this should probably be pushed into a helper function
	sessionCookie, err := r.Cookie(auth.SessionCookieName)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	session, err := ctx.db.GetSession(sessionCookie.Value)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	account, err := ctx.db.GetAccountByID(session.UserID)
	account.ProjectIds = append(account.ProjectIds, project.ID)

	err = ctx.db.UpdateAccount(account)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(project)
}

func updateProject(ctx apiContext, w http.ResponseWriter, r *http.Request) {
	// TODO: this can be refactored
	var project models.Project
	err := json.NewDecoder(r.Body).Decode(&project)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = ctx.db.UpdateProject(project)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(project)
}

func deleteProject(ctx apiContext, w http.ResponseWriter, r *http.Request) {
	var project models.Project
	err := json.NewDecoder(r.Body).Decode(&project)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = ctx.db.RemoveProject(project.ID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}
