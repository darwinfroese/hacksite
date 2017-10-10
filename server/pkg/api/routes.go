package api

import (
	"net/http"
)

const (
	apiPrefix = "/api/v1"
)

var readWriteUpdateMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
var readWriteMethods = []string{"GET", "POST", "OPTIONS"}
var readMethods = []string{"GET", "OPTIONS"}
var writeMethods = []string{"POST", "OPTIONS"}

// TODO: Implement logging

// RegisterRoutes registers all the routes into the mux
func RegisterRoutes(ctx *Context, mux *http.ServeMux) {
	mux.HandleFunc(apiPrefix+"/projects", ctx.projectsRoute)
	mux.HandleFunc(apiPrefix+"/project", ctx.projectRoute)
	mux.HandleFunc(apiPrefix+"/tasks", ctx.tasksRoute)
	mux.HandleFunc(apiPrefix+"/iteration", ctx.iterationsRoute)
	mux.HandleFunc(apiPrefix+"/currentiteration", ctx.currentIterationRoute)
	mux.HandleFunc(apiPrefix+"/accounts", ctx.accountsRoute)
	mux.HandleFunc(apiPrefix+"/login", ctx.loginRoute)
	mux.HandleFunc(apiPrefix+"/logout", ctx.logoutRoute)
	mux.HandleFunc(apiPrefix+"/session", ctx.sessionRoute)
	// TODO: Non-API routes should register somewhere else
	mux.Handle("/health", http.HandlerFunc(healthCheckHandler))
}
