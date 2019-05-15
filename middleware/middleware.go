package middleware

import (
	"log"
	"net/http"
)

// Connected is true when a Supervisor detected downstrea; otherwise false
type Connected bool

// IsConnected is a middleware method to enable graceful handling of requests while no downstream
// Supervisor node is connected to this node
func (c *Connected) IsConnected(handlr http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if *c {
			handlr.ServeHTTP(w, r)
		} else {
			log.Printf("Requeuest for %v received while no peer connection", r.URL.Path)
			http.Error(w, "Herdius network currently undergoing maintenance; we were unable to pass your request to the blockchain network at this time", http.StatusPreconditionFailed)
		}
	})
}
