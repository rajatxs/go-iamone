package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Active        bool               `bson:"active"`
	Username      string             `bson:"username"`
	Fullname      string             `bson:"fullname,omitempty"`
	Email         string             `bson:"email"`
	Bio           string             `bson:"bio,omitempty"`
	Location      string             `bson:"location,omitempty"`
	PasswordHash  string             `bson:"passwordHash"`
	EmailVerified bool               `bson:"emailVerified"`
	ImageUrl      string             `bson:"imageUrl,omitempty"`
	CreatedAt     time.Time          `bson:"createdAt"`
	UpdatedAt     time.Time          `bson:"updatedAt"`
}
