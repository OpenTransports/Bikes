package main

import (
	"os"

	"github.com/OpenTransports/lib-go/models"
	"github.com/go-siris/siris"
)

const citybikesServerURL = "http://api.citybik.es"

var serverURL = os.Getenv("SERVER_URL")

var agencies []models.Agency

func main() {
	var err error
	agencies, err = fetchAgencies()

	if err != nil {
		panic(err)
	}

	app := siris.New()

	// Serve medias files
	app.StaticWeb("/medias", "./medias")

	// Build api
	// /api/transports?latitude=...&longitude=...
	app.Get("/transports", getTransports)
	// /api/agencies?latitude=...&longitude=...
	app.Get("/agencies", getAgencies)

	// Set listening port
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Run server
	app.Run(siris.Addr(":"+port), siris.WithCharset("UTF-8"))
}
