package handlers

import (
	"errors"
	"net/http"
	"second/pkg/helpers"
	"second/pkg/models"
)

var ErrorLoginUsed = errors.New("the login is already used")

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
		YearOfBirth: r.FormValue("b_day"),
		Address:     r.FormValue("address"),
		Number:      r.FormValue("tel_number"),
		ExtraInfo:   r.FormValue("extra_info"),
	}
	extraInfo := models.ExtraInfoForPupilRegistration{
		SchoolName: r.FormValue("school_name"),
		TeacherBIO: models.TeacherBIO{
			Name:       r.FormValue("teacher_name"),
			Surname:    r.FormValue("teacher_surname"),
			Patronymic: r.FormValue("teacher_patronymic"),
		},
		ClassLit: models.ClassLit{
			Number:  r.FormValue("class_num"),
			Literal: r.FormValue("class_lit"),
		},
	}

	auth := models.BasicAuth{
		Login:    pupil.Login,
		Password: pupil.Password,
	}

	err = h.Service.ValidateLoginAndPass(&auth)
	if err != nil {
		helpers.BadRequest(w, h.Logger, err)
		return
	}

	err = h.Service.RegistrationPupil(&pupil, &extraInfo)
	if err != nil {
		if errors.As(err, &ErrorLoginUsed) {
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
