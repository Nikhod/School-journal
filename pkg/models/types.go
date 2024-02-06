package models

type BasicAuth struct {
	Login, Password string
}

type Answer struct {
	Answer string `json:"answer"`
}
