package middleware

import (
	"net/http"
)

type Connected bool

func (c *Connected) IsConnected(handlr http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if *c {
			handlr.ServeHTTP(w, r)
		} else {
			http.Error(w, "Herdius network currently undergoing maintenance; we were unable to pass your request to the blockchain network at this time", http.StatusPreconditionFailed)
		}
	})
}
