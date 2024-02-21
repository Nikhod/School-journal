package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"second/internal/handlers"
)

func InitMux(h *handlers.Handler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/registration_pupil", h.RegistrationPupil).Methods(http.MethodPost)
	router.HandleFunc("/registration_teacher", h.RegistrationTeacher).Methods(http.MethodPost)
	return router
}
