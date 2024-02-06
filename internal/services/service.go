package services

import (
	"errors"
	"go.uber.org/zap"
	"second/internal/repositories"
	"second/pkg/models"
)

type Services struct {
	Repository *repositories.Repository
	Logger     *zap.Logger
}

func NewService(repository *repositories.Repository, logger *zap.Logger) *Services {
	return &Services{Repository: repository, Logger: logger}
}

func (s *Services) ValidateLoginAndPass(auth *models.BasicAuth) error {
	if len(auth.Login) < 8 || len(auth.Login) > 20 {
		return errors.New("Invalid data for registration ")
	}
	if len(auth.Password) < 6 || len(auth.Password) > 20 {
		return errors.New("Invalid data for registration ")
	}
	return nil
}

func (s *Services) RegistrationUser(auth *models.BasicAuth) error {
	isFree := s.Repository.IsLoginFree(auth.Login)
	if isFree == false {
		return errors.New("the login is already used")
	}

	err := s.Repository.RegistrationUser(auth)
	if err != nil {
		return err
	}

	return nil
}
