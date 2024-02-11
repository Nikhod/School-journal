package models

import (
	"time"
)

type Server struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type Db struct {
	DbName   string `json:"db_name"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

type Configs struct {
	Database Db     `json:"database"`
	Server   Server `json:"server"`
}

type Pupil struct {
	Id               uint64    `gorm:"primary_key;auto_increment"`
	Active           bool      `gorm:"default:true"`
	Name             string    `gorm:"type:text;not null"`
	Surname          string    `gorm:"type:text;not null"`
	Patronymic       string    `gorm:"type:text;default:'no data'"`
	Login            string    `gorm:"type:text;not null"`
	Password         string    `gorm:"type:text;not null"`
	YearOfBirth      string    `gorm:"type:text;not null"`
	Address          string    `gorm:"type:text;default:'no data'"`
	Number           string    `gorm:"type:text;default:'no data'"`
	SchoolId         uint      `gorm:"not null"`
	School           School    `gorm:"not null;foreignkey:SchoolId"`
	ExtraInfo        string    `gorm:"type:text;default:'no data'"`
	ClassroomTeacher uint      `gorm:"not null;default:0"`
	ClassLiteral     uint      `gorm:"not null"`
	Teacher          Teacher   `gorm:"foreignkey:ClassroomTeacher"`
	Class            Class     `gorm:"foreignkey:ClassLiteral"`
	CreatedAt        time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt        time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt        time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}

type Class struct {
	Id        uint64    `gorm:"not null;primary key; autoincrement"`
	Active    bool      `gorm:"not null;default:true"`
	Number    int       `gorm:"not null"`
	Literal   string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null;default:current_timestamp"`
	UpdatedAt time.Time `gorm:"not null;default:current_timestamp"`
	DeletedAt time.Time `gorm:"not null;default:current_timestamp"`
}

type Teacher struct {
	Id                  uint64    `gorm:"not null;primary key; autoincrement"`
	Active              bool      `gorm:"not null;default:true"`
	Name                string    `gorm:"not null"`
	Surname             string    `gorm:"not null"`
	Patronymic          string    `gorm:"not null;default:'no data'"`
	YearOfBirth         string    `gorm:"not null"`
	Login               string    `gorm:"type:text;not null"`
	Password            string    `gorm:"type:text;not null"`
	Number              string    `gorm:"not null;default:'no data'"`
	Classroom           string    `gorm:"not null;default:'no data'"`
	ClassroomManagement uint      `gorm:"not null;default:0"`
	Class               Class     `gorm:"not null;foreignkey:ClassroomManagement"`
	SchoolId            uint      `gorm:"not null"`
	School              School    `gorm:"not null; foreignkey:SchoolId"`
	CreatedAt           time.Time `gorm:"not null;default:current_timestamp"`
	UpdatedAt           time.Time `gorm:"not null;default:current_timestamp"`
	DeletedAt           time.Time `gorm:"not null;default:current_timestamp"`
}

type Subject struct {
	Id            uint64      `gorm:"not null;primary key; autoincrement"`
	Active        bool        `gorm:"not null;default:true"`
	CreatedAt     time.Time   `gorm:"not null;default:current_timestamp"`
	UpdatedAt     time.Time   `gorm:"not null;default:current_timestamp"`
	DeletedAt     time.Time   `gorm:"not null;default:current_timestamp"`
	SubjectNameId int         `gorm:"not null"`
	TeacherId     int         `gorm:"not null"`
	SubjectName   SubjectName `gorm:"not null;foreignkey:SubjectNameId"`
	Teacher       Teacher     `gorm:"not null;foreignkey:TeacherId"`
}

type Circle struct {
	Id        uint64    `gorm:"not null;primary key; autoincrement"`
	Active    bool      `gorm:"not null;default:true"`
	PupilId   int       `gorm:"not null"`
	SubjectId int       `gorm:"not null"`
	TeacherId int       `gorm:"not null;"`
	Teacher   Teacher   `gorm:"not null;foreignkey:TeacherId"`
	Pupil     Pupil     `gorm:"not null;foreignkey:PupilId"`
	Subject   Subject   `gorm:"not null;foreignkey:SubjectId"`
	CreatedAt time.Time `gorm:"not null;default:current_timestamp"`
	UpdatedAt time.Time `gorm:"not null;default:current_timestamp"`
	DeletedAt time.Time `gorm:"not null;default:current_timestamp"`
}

type SubjectName struct {
	Id          uint64    `gorm:"not null;primary key; autoincrement"`
	Active      bool      `gorm:"not null;default:true"`
	SubjectName string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"not null;default:current_timestamp"`
	UpdatedAt   time.Time `gorm:"not null;default:current_timestamp"`
	DeletedAt   time.Time `gorm:"not null;default:current_timestamp"`
}

type School struct {
	Id         uint64    `gorm:"not null;primary key; autoincrement"`
	Active     bool      `gorm:"not null;default:true"`
	SchoolName string    `gorm:"type:text; not null"`
	CreatedAt  time.Time `gorm:"not null;default:current_timestamp"`
	UpdatedAt  time.Time `gorm:"not null;default:current_timestamp"`
	DeletedAt  time.Time `gorm:"not null;default:current_timestamp"`
}
