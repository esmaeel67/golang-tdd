package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func ConfigureServer(handler *Handler) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("GET").Path("/").Handler(http.HandlerFunc(handler.Index))

	return router
}
