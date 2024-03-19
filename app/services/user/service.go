package user

import (
	"fiber-boilerplate/app/models"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterInput) (models.User, error)
	Login(input LoginInput) (models.User, error)
}

type service struct {
	repo Repository
}

func CreateService(repo Repository) Service {
	return &service{repo}
}

func (s *service) RegisterUser(input RegisterInput) (models.User, error) {
	user := models.User{}
	user.Email = input.Email
	user.Username = input.Username

	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}

	user.Password = string(password)
	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.CreditBalance = 10 //default

	newUser, err := s.repo.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Login(input LoginInput) (models.User, error) {
	user, err := s.repo.FindByEmailOrUsername(input.Query)
	if err != nil {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return user, err
	}

	return user, nil
}
