package handlers

import (
	"errors"
	"net/http"
	"second/pkg/helpers"
	"second/pkg/models"
)

func (h *Handler) Registration(w http.ResponseWriter, r *http.Request) {
	auth := models.BasicAuth{
		Login:    r.Header.Get("login"),
		Password: r.Header.Get("pass"),
	}

	err := h.Service.ValidateLoginAndPass(&auth)
	if err != nil {
		helpers.BadRequest(w, h.Logger, err)
		return
	}

	err = h.Service.RegistrationUser(&auth)
	if err != nil {
		if errors.Is(err, errors.New("the login is already used")) {
			helpers.Forbidden(w, h.Logger, err)
			return
		}
		helpers.InternalServerError(w, h.Logger, err)
		return
	}

	err = helpers.SendAnswer(w, "Регистрация прошла успешно")
	if err != nil {
		helpers.BadRequest(w, h.Logger, err)
		return
	}
}

func (h *Handler) GetToken(w http.ResponseWriter, r *http.Request) {
	
}