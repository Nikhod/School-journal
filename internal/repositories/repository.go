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
