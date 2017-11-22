package api

import (
	"net/http"
)

const (
	apiPrefix = "/api/v1"
)

// RegisterAPIRoutes registers all the api routes into the mux
func RegisterAPIRoutes(ctx *Context, mux *http.ServeMux) {
	mux.HandleFunc(apiPrefix+"/projects", Apply(ctx, ctx.projectsRoute))
	mux.HandleFunc(apiPrefix+"/project", Apply(ctx, ctx.projectRoute))
	mux.HandleFunc(apiPrefix+"/tasks", Apply(ctx, ctx.tasksRoute))
	mux.HandleFunc(apiPrefix+"/evolution", Apply(ctx, ctx.evolutionsRoute))
	mux.HandleFunc(apiPrefix+"/currentevolution", Apply(ctx, ctx.currentEvolutionRoute))
	mux.HandleFunc(apiPrefix+"/accounts", Apply(ctx, ctx.accountsRoute))
	mux.HandleFunc(apiPrefix+"/login", Apply(ctx, ctx.loginRoute))
	mux.HandleFunc(apiPrefix+"/logout", Apply(ctx, ctx.logoutRoute))
	mux.HandleFunc(apiPrefix+"/session", Apply(ctx, ctx.sessionRoute))
}

// RegisterRoutes registers all non api routes into the mux
func RegisterRoutes(ctx *Context, mux *http.ServeMux) {
	mux.Handle("/health", Apply(ctx, http.HandlerFunc(healthCheckHandler)))
}
