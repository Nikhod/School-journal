package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"second/internal/handlers"
)

func InitMux(h *handlers.Handler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/registration", h.RegistrationPupil).Methods(http.MethodPost)
	return router
}
