package handler

import (
	"log"
	"net/http"
)

type Connected bool

func (c *Connected) IsConnected(handlr http.Handler) http.Handler {
	log.Println("in middlware")
	log.Println("c:", c)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if *c {
			handlr.ServeHTTP(w, r)
		} else {
			http.Error(w, "unable to pass request to Supervisor", http.StatusForbidden)
		}
	})
}
