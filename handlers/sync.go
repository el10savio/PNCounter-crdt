package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/el10savio/pncounter-crdt/pncounter"
	log "github.com/sirupsen/logrus"
)

// Sync merges multiple PNCounters present in a network to get them in sync
// It does so by obtaining the PNCounter from each node in the cluster
// and performs a merge operation with the local PNCounter
func Sync(PNCounter pncounter.PNCounter) (pncounter.PNCounter, error) {
	// Obtain addresses of peer nodes in the cluster
	peers := GetPeerList()

	// Return the local PNCounter back if no peers
	// are present along with an error
	if len(peers) == 0 {
		return PNCounter, errors.New("nil peers present")
	}

	// Iterate over the peer list and send a /pncounter/values GET request
	// to each peer to obtain its PNCounter
	for _, peer := range peers {
		peerPNCounter, err := SendListRequest(peer)
		if err != nil {
			log.WithFields(log.Fields{"error": err, "peer": peer}).Error("failed sending pncounter values request")
			continue
		}

		// Merge the peer's PNCounter with our local PNCounter
		PNCounter = pncounter.Merge(PNCounter, peerPNCounter)
	}

	// DEBUG log in the case of success
	// indicating the new PNCounter
	log.WithFields(log.Fields{
		"count": PNCounter,
	}).Debug("successful pncounter sync")

	// Return the synced new PNCounter
	return PNCounter, nil
}

// SendListRequest is used to send a GET /pncounter/values
// to peer nodes in the cluster
func SendListRequest(peer string) (pncounter.PNCounter, error) {
	var _pncounter pncounter.PNCounter

	// Return an empty PNCounter followed by an error if the peer is nil
	if peer == "" {
		return _pncounter, errors.New("empty peer provided")
	}

	// Resolve the Peer ID and network to generate the request URL
	url := fmt.Sprintf("http://%s.%s/pncounter/values", peer, GetNetwork())
	response, err := SendRequest(url)
	if err != nil {
		return _pncounter, err
	}

	// Return an empty PNCounter followed by an error
	// if the peer's response is not HTTP 200 OK
	if response.StatusCode != http.StatusOK {
		return _pncounter, errors.New("received invalid http response status:" + fmt.Sprint((response.StatusCode)))
	}

	// Decode the peer's PNCounter to be usable by our local PNCounter
	var pnCounter pncounter.PNCounter
	err = json.NewDecoder(response.Body).Decode(&pnCounter)
	if err != nil {
		return _pncounter, err
	}

	// Return the decoded peer's PNCounter
	_pncounter = pnCounter
	return _pncounter, nil
}
