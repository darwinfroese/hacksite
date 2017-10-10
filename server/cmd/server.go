package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/darwinfroese/hacksite/server/pkg/api"
	"github.com/darwinfroese/hacksite/server/pkg/config"
	"github.com/darwinfroese/hacksite/server/pkg/database"
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
	fmt.Println("Setting up the server.")

	// Redirect *:80 to *:443
	go http.ListenAndServe(":80", http.HandlerFunc(api.RedirectToHTTPS))

	m := http.NewServeMux()
	db := database.CreateDB()
	c := config.ParseConfig(envFile)

	ctx := &api.Context{DB: &db, Config: &c}

	m.Handle("/", http.FileServer(http.Dir(c.WebFileLocation)))

	api.RegisterRoutes(ctx, m)

	fmt.Println("Starting server scheduler.")
	scheduler.Start(ctx)

	if _, err := os.Stat(c.CertLocation); os.IsNotExist(err) {
		fmt.Println("Couldn't find: ", c.CertLocation)
	}

	if _, err := os.Stat(c.KeyLocation); os.IsNotExist(err) {
		fmt.Println("Couldn't find: ", c.KeyLocation)
	}

	fmt.Println("Starting the server.")
	fmt.Println("Server failed with: ", http.ListenAndServeTLS(
		c.Port,
		c.CertLocation,
		c.KeyLocation,
		m))
}
