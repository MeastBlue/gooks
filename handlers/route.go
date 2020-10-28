package handlers

import (
	"books/models"

	"github.com/gorilla/mux"
)

func NewRouter(routes models.Routes) *mux.Router {
	router := mux.NewRouter()
	router.StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.Func)
	}

	return router
}
