package repositories

import (
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
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
	err := r.Database.Create(&pupil).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) IsLoginFree(login string) bool {
	var pupil models.Pupil
	var amountOfRecords int64
	tx := r.Database.Select("id").First(&pupil, models.Pupil{Login: login}).Count(&amountOfRecords)

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
	var school models.School
	tx := r.Database.
		Model(models.School{SchoolName: schoolName}).
		Select("id").
		Where(models.School{SchoolName: schoolName, Active: true}).
		Count(&amountOfRecords)

	if amountOfRecords == 0 {
		return 0, gorm.ErrRecordNotFound
	}

	err = tx.First(&school).Error
	if err != nil {
		return 0, err
	}

	schoolID = school.Id
	return schoolID, nil
}

func (r *Repository) GetTeacherIDByByBIO(extraInfo *models.ExtraInfoForRegistration) (teacherID uint, err error) {
	var amountOfRecords int64
	var teacher models.Teacher
	tx := r.Database.
		Model(models.Teacher{}).
		Select("id").
		Where(models.Teacher{Name: extraInfo.TeacherBIO.Name, Surname: extraInfo.TeacherBIO.Surname,
			Patronymic: extraInfo.TeacherBIO.Patronymic, Active: true}).
		Count(&amountOfRecords)

	if amountOfRecords == 0 {
		return 0, gorm.ErrRecordNotFound
	}

	err = tx.First(&teacher).Error
	if err != nil {
		return 0, err
	}

	teacherID = teacher.Id
	return teacherID, nil
}

func (r *Repository) GetClassID(classLit *models.ClassLit) (classID uint, err error) {
	var amountOfRecord int64
	var class models.Class
	tx := r.Database.
		Model(models.Class{}).
		Select("id").
		Where(models.Class{Number: classLit.Number, Literal: classLit.Literal, Active: true}).
		Count(&amountOfRecord)

	if amountOfRecord == 0 {
		return 0, gorm.ErrRecordNotFound
	}

	err = tx.First(&class).Error
	if err != nil {
		return 0, err
	}

	classID = class.Id
	return classID, nil
}

func (r *Repository) RegistrationTeacher(teacher *models.Teacher) error {
	err := r.Database.Create(&teacher).Error
	if err != nil {
		return err
	}
	return nil
}
