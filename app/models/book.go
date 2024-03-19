package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Owner       User               `bson:"owner,omitempty" json:"owner"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at,omitempty" json:"updated_at"`
	IsDeleted   bool               `bson:"is_deleted" json:"is_deleted"`
	Title       string             `bson:"title" json:"title" validate:"required,lte=255"`
	Description string             `bson:"description,omitempty" json:"description"`
	Picture     string             `bson:"picture,omitempty" json:"picture"`
	Rating      int                `bson:"rating" json:"rating" validate:"min=1,max=10"`
}

func (b Book) Validate() error {
	return validator.New().Struct(b)
}
