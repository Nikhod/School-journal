package repositories

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"log"
	"second/pkg/database"
	"second/pkg/models"
)

type Repository struct {
	Database *gorm.DB
	Logger   *zap.Logger
}

func NewRepository(database *gorm.DB, logger *zap.Logger) *Repository {
	return &Repository{Database: database, Logger: logger}
}

func (r *Repository) RegistrationUser(auth *models.BasicAuth) error {
	var user database.User
	if err := r.Database.First(&user, 2).Error; err != nil {
		log.Println("data not found")
		return err
	}

	err := r.Database.Model(&user).Association("role").Find(&user.Role)
	if err != nil {
		log.Println(err)
		log.Println("error here")
		return err
	}

	log.Println(user)
	fmt.Println()
	log.Println(user.Role.Role)
	return nil
}

func (r *Repository) IsLoginFree(login string) bool {
	//todo ligic
	return true
}
