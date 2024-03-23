package user

import (
	"errors"
	"fiber-boilerplate/app/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
	user, err := s.repo.FindByEmailOrUsername(input.Email)
	if err != nil {
		return models.User{}, err
	}

	if user.ID != primitive.NilObjectID {
		return models.User{}, errors.New("email already exists")
	}

	user, err = s.repo.FindByEmailOrUsername(input.Username)
	if err != nil {
		return models.User{}, err
	}
	if user.ID != primitive.NilObjectID {
		return models.User{}, errors.New("username already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	newUser := models.User{
		Email:         input.Email,
		Username:      input.Username,
		Password:      string(hashedPassword),
		FirstName:     input.FirstName,
		LastName:      input.LastName,
		CreditBalance: 10, //default
	}

	insertedID, err := s.repo.Save(newUser)
	if err != nil {
		return models.User{}, err
	}

	insertedUser, err := s.repo.FindByID(insertedID)
	if err != nil {
		return models.User{}, err
	}

	return insertedUser, nil
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
