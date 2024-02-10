package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"second/pkg/models"
)

func InitDatabase(configs *models.Configs) (*gorm.DB, error) {
	dbURI := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		configs.Database.Host, configs.Database.User, configs.Database.Password,
		configs.Database.DbName, configs.Database.Port)

	database, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = Migrations(database)
	if err != nil {
		return nil, err
	}

	return database, nil
}

func Migrations(db *gorm.DB) error {

	migrator := db.Migrator()

	err := migrator.AutoMigrate(&models.SubjectName{}, &models.Class{},
		&models.Teacher{}, &models.Subject{}, &models.Pupil{}, &models.Circle{})
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
