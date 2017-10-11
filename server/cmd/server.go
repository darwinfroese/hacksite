package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/darwinfroese/hacksite/server/pkg/api"
	"github.com/darwinfroese/hacksite/server/pkg/config"
	"github.com/darwinfroese/hacksite/server/pkg/database/bolt"
	"github.com/darwinfroese/hacksite/server/pkg/log/logrus"
	"github.com/darwinfroese/hacksite/server/pkg/scheduler"
)

// These are vars so they can be set at compile time
var (
	version = "1.0.0"
	// give this a fallback
	envFile = "environments/dev.env.json"
)

// TODO: Receive arguments to perform actions on the server while it's running
// TODO: Receive arguments to gracefully shutdown the server

func main() {

	// Setup
	m := http.NewServeMux()
	db := bolt.New()
	c := config.ParseConfig(envFile)
	logger := logrus.New(c.Logger.LogFileLocation)

	ctx := api.Context{DB: &db, Config: &c, Logger: &logger}

	logger.Info("Starting redirect server")
	// Redirect *:80 to *:443
	go http.ListenAndServe(":80", http.HandlerFunc(api.RedirectToHTTPS))
	m.Handle("/", http.FileServer(http.Dir(c.Server.WebFileLocation)))

	api.RegisterRoutes(m)
	api.RegisterAPIRoutes(&ctx, m)

	logger.Info("Starting scheduler")
	scheduler.Start(ctx)

	if _, err := os.Stat(c.Server.CertLocation); os.IsNotExist(err) {
		logger.Error(fmt.Sprintf("Couldn't find %s", c.Server.CertLocation))
		return
	}

	if _, err := os.Stat(c.Server.KeyLocation); os.IsNotExist(err) {
		logger.Error(fmt.Sprintf("Couldn't find %s", c.Server.KeyLocation))
		return
	}

	logger.Info("Starting api server.")
	logger.Error(fmt.Sprintf("Server failed with %s", http.ListenAndServeTLS(
		c.Server.Port,
		c.Server.CertLocation,
		c.Server.KeyLocation,
		m)))
}
