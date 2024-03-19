package book

import (
	"fiber-boilerplate/app/models"
)

type Service interface {
	CreateBook(input CreateBookInput, loggedInUser models.User) (models.Book, error)
}

type service struct {
	repo Repository
}

func CreateService(repo Repository) Service {
	return &service{repo}
}

func (s *service) CreateBook(input CreateBookInput, loggedInUser models.User) (models.Book, error) {
	book := models.Book{}
	book.Title = input.Title
	book.Description = input.Description
	book.Owner = loggedInUser

	newBook, err := s.repo.Save(book)
	if err != nil {
		return newBook, err
	}

	return newBook, nil
}
