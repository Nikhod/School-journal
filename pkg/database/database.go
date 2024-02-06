package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"second/pkg/models"
	"time"
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

type Role struct {
	ID        uint      `gorm:"not null;primary_key;autoincrement"`
	Role      string    `gorm:"type:text;not null;default:client"`
	Active    bool      `gorm:"not null;default:true"`
	CreatedAt time.Time `gorm:"type:timestamptz;not null;default:current_timestamp"`
	UpdatedAt time.Time `gorm:"type:timestamptz;not null;default:current_timestamp"`
	DeletedAt time.Time `gorm:"type:timestamptz;not null;default:current_timestamp"`
	Something int       `gorm:"type:bigint;not null;default:0"`
}
type User struct {
	ID        uint      `gorm:"not null;primary_key;autoincrement"`
	Age       int       `gorm:"not null;default:2000"`
	Name      string    `gorm:"not null;default:no data"`
	Login     string    `gorm:"not null;default:no data"`
	Password  string    `gorm:"not null;default:no data"`
	RoleID    uint      `gorm:"not null;default:23"`
	Role      Role      `gorm:"not null;foreignkey:RoleID"`
	Active    bool      `gorm:"not null;default:true"`
	CreatedAt time.Time `gorm:"not null;default:current_timestamp"`
	UpdatedAt time.Time `gorm:"not null;default:current_timestamp"`
	DeletedAt time.Time `gorm:"not null;default:current_timestamp"`
	Anything  string    `gorm:"type:text;not null;default:no data"`
}
