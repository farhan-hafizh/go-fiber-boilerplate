package book

import (
	"context"
	"go-fiber-boilerplate/app/models"
	"go-fiber-boilerplate/database"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Save(book models.Book) (models.Book, error)
}

type repository struct {
	collection *mongo.Collection
}

func CreateRepository(db *database.DB) Repository {
	collection := db.Collection("books")
	return &repository{collection}
}

func (repo *repository) Save(book models.Book) (models.Book, error) {
	if err := book.Validate(); err != nil {
		return book, err
	}

	_, err := repo.collection.InsertOne(context.Background(), book)
	if err != nil {
		return book, err
	}

	return book, nil
}
