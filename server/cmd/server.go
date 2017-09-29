package main

import (
	"fmt"
	"net/http"

	"github.com/darwinfroese/hacksite/server/pkg/api"
	"github.com/darwinfroese/hacksite/server/pkg/database"
	"github.com/darwinfroese/hacksite/server/pkg/scheduler"
)

// Constants
const (
	version = "1.0.0"
)

// TODO: Receive arguments to perform actions on the server while it's running
// TODO: Receive arguments to gracefully shutdown the server

func main() {
	fmt.Println("Setting up the server.")

	m := http.NewServeMux()
	db := database.CreateDB()

	//m.Handle("/", http.FileServer(http.Dir("./webdist")))
	m.Handle("/", http.FileServer(http.Dir("/var/www/hacksite")))

	api.RegisterRoutes(m, db)

	fmt.Println("Starting server scheduler.")
	scheduler.Start(db)

	fmt.Println("Starting the server.")
	http.ListenAndServe(":80", m)
}
