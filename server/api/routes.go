package api

import (
	"net/http"

	"github.com/darwinfroese/hacksite/server/database"
)

const (
	apiPrefix = "/api/v1"
)

// RegisterRoutes registers all the routes into the mux
func RegisterRoutes(mux *http.ServeMux, db database.Database) {
	mux.Handle(apiPrefix+"/projects", context{db: db, handler: projects})
	mux.Handle(apiPrefix+"/project", context{db: db, handler: getProject})
	mux.Handle(apiPrefix+"/tasks", context{db: db, handler: tasks})
}
