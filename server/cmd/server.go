package main

import (
	"fmt"
	"net/http"

	"github.com/darwinfroese/hacksite/server/pkg/api"
	"github.com/darwinfroese/hacksite/server/pkg/config"
	"github.com/darwinfroese/hacksite/server/pkg/database"
	"github.com/darwinfroese/hacksite/server/pkg/scheduler"
)

// Constants
const (
	version = "1.0.0"
	envFile = "dev.env.json"
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
	m.Handle("/", http.FileServer(http.Dir(c.WebFileLocation)))

	api.RegisterRoutes(m, db)

	fmt.Println("Starting server scheduler.")
	scheduler.Start(db)

	fmt.Println("Starting the server.")
	http.ListenAndServeTLS(
		c.Port,
		c.CertLocation,
		c.KeyLocation,
		m,
	)
}
