package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	Email        *string            `json:"email" validate:"required, email"`
	Password     *string            `json:"-" validate:"required, min=6, max=100"`
	Name         *string            `json:"name" validate:"required, min=1, max=100"`
	ImageURL     *string            `json:"imageUrl"`
	Token        *string            `json:"token"`
	UserType     *string            `json:"userType" validate:"required, eq=USER|eq=ADMIN"`
	RefreshToken *string            `json:"refreshToken"`
	Website      *string            `json:"website"`
	CreatedAt    time.Time          `json:"createdAt"`
	UpdatedAt    time.Time          `json:"updatedAt"`
	UserID       *string            `json:"userId"`
}
