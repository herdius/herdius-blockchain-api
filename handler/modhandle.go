package handler

import (
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

type RouterSwapper struct {
	Mu   sync.Mutex
	Root *mux.Router
}

func (rs *RouterSwapper) Swap(newRouter *mux.Router) {
	rs.Mu.Lock()
	rs.Root = newRouter
	rs.Mu.Unlock()
}

func (rs RouterSwapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rs.Mu.Lock()
	root := rs.Root
	rs.Mu.Unlock()
	root.ServeHTTP(w, r)
}
