package main

// The following implements the main Go
// package starting up the pncounter server

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/el10savio/pncounter-crdt/handlers"
)

const (
	// PORT used for the PNCounter server
	PORT = "8080"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	r := handlers.Router()

	log.WithFields(log.Fields{
		"port": PORT,
	}).Info("started PNCounter node server")

	http.ListenAndServe(":"+PORT, r)
}
