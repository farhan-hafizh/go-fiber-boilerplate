package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email         string             `bson:"email" json:"email" validate:"required,email"`
	Username      string             `bson:"username" json:"username" validate:"required,lte=255"`
	Password      string             `bson:"password" json:"password" validate:"required,lte=255"`
	Photo         string             `bson:"photo" json:"photo"`
	FirstName     string             `bson:"first_name" json:"first_name" validate:"required,lte=255"`
	LastName      string             `bson:"last_name" json:"last_name" validate:"required,lte=255"`
	CreditBalance int32              `bson:"credit_balance" json:"credit_balance"`
	CreatedAt     time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at,omitempty" json:"updated_at"`
}

func (u User) Validate() error {
	return validator.New().Struct(u)
}
