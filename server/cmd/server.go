package main

import (
	"fmt"
	"net/http"

	"github.com/darwinfroese/hacksite/server/api"
	"github.com/darwinfroese/hacksite/server/database"
)

const (
	version = "1.0.0"
)

func main() {
	fmt.Println("Setting up the server.")

	m := http.NewServeMux()
	db := database.CreateBoltDB()

	m.Handle("/", http.FileServer(http.Dir("./webdist")))

	api.RegisterRoutes(m, db)

	fmt.Println("Starting the server.")
	http.ListenAndServe(":8800", m)
}
