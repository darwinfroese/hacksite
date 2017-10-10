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

// RegisterAPIRoutes registers all the api routes into the mux
func RegisterAPIRoutes(ctx Context, mux *http.ServeMux) {
	mux.HandleFunc(apiPrefix+"/projects", Apply(ctx, ctx.projectsRoute))
	mux.HandleFunc(apiPrefix+"/project", Apply(ctx, ctx.projectRoute))
	mux.HandleFunc(apiPrefix+"/tasks", Apply(ctx, ctx.tasksRoute))
	mux.HandleFunc(apiPrefix+"/iteration", Apply(ctx, ctx.iterationsRoute))
	mux.HandleFunc(apiPrefix+"/currentiteration", Apply(ctx, ctx.currentIterationRoute))
	mux.HandleFunc(apiPrefix+"/accounts", Apply(ctx, ctx.accountsRoute))
	mux.HandleFunc(apiPrefix+"/login", Apply(ctx, ctx.loginRoute))
	mux.HandleFunc(apiPrefix+"/logout", Apply(ctx, ctx.logoutRoute))
	mux.HandleFunc(apiPrefix+"/session", Apply(ctx, ctx.sessionRoute))
}

// RegisterRoutes registers all non api routes into the mux
func RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("/health", setCorsHeaders(Context{}, http.HandlerFunc(healthCheckHandler)))
}
