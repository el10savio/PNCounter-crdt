package handlers

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Decrement is the HTTP handler used to decrement
// the count of the PNCounter node in the server
func Decrement(w http.ResponseWriter, r *http.Request) {
	// Decrement the given value to our stored PNCounter
	PNCounter = PNCounter.Decrement(GetMyNodeIP())

	// DEBUG log in the case of success indicating
	// the new PNCounter count
	log.WithFields(log.Fields{
		"count": PNCounter,
	}).Debug("successful pncounter decrement")

	// Return HTTP 200 OK in the case of success
	w.WriteHeader(http.StatusOK)
}
