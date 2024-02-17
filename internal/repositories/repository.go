package repositories

import (
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"log"
	"second/pkg/models"
)

type Repository struct {
	Database *gorm.DB
	Logger   *zap.Logger
}

func NewRepository(database *gorm.DB, logger *zap.Logger) *Repository {
	return &Repository{Database: database, Logger: logger}
}

func (r *Repository) RegistrationPupil(pupil *models.Pupil) error {
	r.Database.Create(pupil)
	return nil
}

func (r *Repository) IsLoginFree(login string) bool {
	var pupil models.Pupil
	var amountOfRecords int64
	tx := r.Database.Select("id").First(&pupil, models.Pupil{Login: login}).Count(&amountOfRecords)
	log.Println(amountOfRecords)

	err := tx.Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return true
	} else if err != nil {
		r.Logger.Error("error during the registration")
		return false
	}
	if amountOfRecords != 0 {
		return false
	}

	return true
}

func (r *Repository) GetSchoolIDByName(schoolName string) (schoolID uint, err error) {
	var amountOfRecords int64
	tx := r.Database.
		Select("id").
		First(&schoolID, models.School{SchoolName: schoolName}).
		Count(&amountOfRecords)

	if err = tx.Error; err != nil {
		return 0, err
	}

	if amountOfRecords == 0 {
		return 0, gorm.ErrRecordNotFound
	}

	return schoolID, nil
}

func (r *Repository) GetTeacherIDByByBIO(extraInfo *models.ExtraInfoForPupilRegistration) (teacherID uint, err error) {
	var amountOfRecords int64
	r.Database.Select("id").
		First(&teacherID, models.Teacher{Name: extraInfo.TeacherBIO.Name, Surname: extraInfo.TeacherBIO.Surname,
			Patronymic: extraInfo.TeacherBIO.Patronymic}).
		Count(&amountOfRecords)

	if amountOfRecords == 0 {
		return 0, gorm.ErrRecordNotFound
	}

	return teacherID, nil
}

func (r *Repository) GetClassID(classLit *models.ClassLit) (classID uint, err error) {
	var amountOfRecord int64

	r.Database.Select("id").
		First(&classID, models.Class{Number: classLit.Number, Literal: classLit.Literal}).
		Count(&amountOfRecord)

	if amountOfRecord == 0 {
		return 0, gorm.ErrRecordNotFound
	}

	return classID, nil
}
