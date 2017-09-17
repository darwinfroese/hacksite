package api

import (
	"net/http"

	"github.com/darwinfroese/hacksite/server/database"
)

const (
	apiPrefix = "/api/v1"
)

// TODO: Implement logging

// RegisterRoutes registers all the routes into the mux
func RegisterRoutes(mux *http.ServeMux, db database.Database) {
	mux.Handle(apiPrefix+"/projects", context{db: db, handler: projects})
	mux.Handle(apiPrefix+"/project", context{db: db, handler: project})
	mux.Handle(apiPrefix+"/tasks", context{db: db, handler: tasks})
	mux.Handle(apiPrefix+"/iteration", context{db: db, handler: iterations})
	mux.Handle(apiPrefix+"/currentiteration", context{db: db, handler: currentIteration})
	mux.Handle(apiPrefix+"/accounts", context{db: db, handler: accounts})
	mux.Handle(apiPrefix+"/login", context{db: db, handler: login})
	mux.Handle(apiPrefix+"/session", context{db: db, handler: session})
}
