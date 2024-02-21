package models

type BasicAuth struct {
	Login, Password string
}

type Answer struct {
	Answer string `json:"answer"`
}

type ExtraInfoForRegistration struct {
	SchoolName string
	TeacherBIO TeacherBIO
	ClassLit   ClassLit
}

type TeacherBIO struct {
	Name, Surname, Patronymic string
}

// ClassLit lit - literal
type ClassLit struct {
	Number, Literal string
}
