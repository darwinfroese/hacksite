package api

import (
	"net/http"

	"github.com/darwinfroese/hacksite/server/database"
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
func RegisterRoutes(mux *http.ServeMux, db database.Database) {
	mux.Handle(apiPrefix+"/projects", apiContext{db: db, apiHandler: projects, supportedMethods: readWriteUpdateMethods})
	mux.Handle(apiPrefix+"/project", apiContext{db: db, apiHandler: project, supportedMethods: readMethods})
	mux.Handle(apiPrefix+"/tasks", apiContext{db: db, apiHandler: tasks, supportedMethods: readWriteUpdateMethods})
	mux.Handle(apiPrefix+"/iteration", apiContext{db: db, apiHandler: iterations, supportedMethods: writeMethods})
	mux.Handle(apiPrefix+"/currentiteration", apiContext{db: db, apiHandler: currentIteration, supportedMethods: writeMethods})
	mux.Handle(apiPrefix+"/accounts", apiContext{db: db, apiHandler: accounts, supportedMethods: writeMethods})
	mux.Handle(apiPrefix+"/login", apiContext{db: db, apiHandler: login, supportedMethods: readMethods})
	mux.Handle(apiPrefix+"/logout", apiContext{db: db, apiHandler: logout, supportedMethods: readMethods})
	mux.Handle(apiPrefix+"/session", apiContext{db: db, apiHandler: session, supportedMethods: readMethods})
}
