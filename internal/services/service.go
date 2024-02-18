package services

import (
	"errors"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
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

func (s *Services) RegistrationPupil(pupil *models.Pupil, extraInfo *models.ExtraInfoForPupilRegistration) error {
	isFree := s.Repository.IsLoginFree(pupil.Login)
	if isFree == false {
		return errors.New("the login is already used")
	}

	schoolID, err := s.GetSchoolIDByName(extraInfo.SchoolName)
	if err != nil {
		return err
	}

	teacherID, err := s.Repository.GetTeacherIDByByBIO(extraInfo)
	if err != nil {
		return err
	}

	classLitID, err := s.Repository.GetClassID(&extraInfo.ClassLit)
	if err != nil {
		return err
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(pupil.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	pupil.SchoolId = schoolID
	pupil.ClassroomTeacher = teacherID
	pupil.ClassLiteral = classLitID
	pupil.Password = string(hashedPass)

	err = s.Repository.RegistrationPupil(pupil)
	if err != nil {
		return err
	}

	return nil
}

func (s *Services) GetSchoolIDByName(schoolName string) (schoolID uint, err error) {
	schoolID, err = s.Repository.GetSchoolIDByName(schoolName)
	if err != nil {
		return 0, err
	}

	return schoolID, nil
}
