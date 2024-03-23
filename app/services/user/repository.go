package user

import (
	"context"
	"fiber-boilerplate/app/models"
	"fiber-boilerplate/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	FindByEmailOrUsername(query string) (user models.User, err error)
	Save(user models.User) (insertedId string, err error)
	FindByID(id string) (models.User, error)
}

type repository struct {
	collection *mongo.Collection
}

func CreateRepository(db *database.DB) Repository {
	collection := db.Collection("users")
	return &repository{collection}
}

func (repo *repository) FindByEmailOrUsername(query string) (user models.User, err error) {
	err = repo.collection.FindOne(context.Background(), bson.M{
		"$or": []bson.M{
			{"email": query},
			{"username": query},
		},
	}).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repo *repository) Save(user models.User) (string, error) {
	if err := user.Validate(); err != nil {
		return "", err
	}
	result, err := repo.collection.InsertOne(context.Background(), user)
	if err != nil {
		return "", err
	}

	// Access the inserted user using the result (assuming InsertedID method)
	insertedID := result.InsertedID.(primitive.ObjectID).Hex()

	return insertedID, nil
}

func (repo *repository) FindByID(id string) (models.User, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.User{}, err
	}

	var user models.User
	err = repo.collection.FindOne(context.Background(), bson.M{"_id": objectId}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
