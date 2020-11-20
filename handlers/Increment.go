package handlers

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Increment is the HTTP handler used to increment
// the count of the PNCounter node in the server
func Increment(w http.ResponseWriter, r *http.Request) {
	// Increment the given value to our stored PNCounter
	PNCounter = PNCounter.Increment(GetMyNodeIP())

	// DEBUG log in the case of success indicating
	// the new PNCounter count
	log.WithFields(log.Fields{
		"count": PNCounter,
	}).Debug("successful pncounter increment")

	// Return HTTP 200 OK in the case of success
	w.WriteHeader(http.StatusOK)
}
