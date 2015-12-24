package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jaaaaaaaaaam/go-todo/logger"
)

// NewRouter is the router
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = logger.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
