package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"second/pkg/helpers"
	"second/pkg/models"
)

func (h *Handler) RegistrationPupil(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.InternalServerError(w, h.Logger, err)
		return
	}

	var pupil = models.Pupil{
		Name:        r.FormValue("name"),
		Surname:     r.FormValue("surname"),
		Patronymic:  r.FormValue("patronymic"),
		Login:       r.FormValue("login"),
		Password:    r.FormValue("password"),
		YearOfBirth: r.FormValue("b-day"),
		Address:     r.FormValue("address"),
		Number:      r.FormValue("number"),
		ExtraInfo:   r.FormValue("extra-info"),
	}
	schoolName := r.FormValue("school_name")

	auth := models.BasicAuth{
		Login:    pupil.Login,
		Password: pupil.Password,
	}

	err = h.Service.ValidateLoginAndPass(&auth)
	if err != nil {
		helpers.BadRequest(w, h.Logger, err)
		return
	}

	err = h.Service.RegistrationPupil(&pupil, schoolName)
	if err != nil {
		if errors.Is(err, errors.New("the login is already used")) {
			helpers.Forbidden(w, h.Logger, err)
			return
		}
		helpers.InternalServerError(w, h.Logger, err)
		return
	}

	err = helpers.SendAnswer(w, "registration has finished successfully")
	if err != nil {
		helpers.BadRequest(w, h.Logger, err)
		return
	}
}

func (h *Handler) GetToken(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
}
