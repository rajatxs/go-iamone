package db

import (
	"context"
	"time"

	"github.com/rajatxs/go-iamone/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Returns reference of user collection
func collection() *mongo.Collection {
	return db.Collection("users")
}

// Returns indexes for user collection
func getIndexes() []mongo.IndexModel {
	// unique constraint on email
	emailIndex := mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	}

	// unique constraint on username
	usernameIndex := mongo.IndexModel{
		Keys:    bson.M{"username": 1},
		Options: options.Index().SetUnique(true),
	}

	return []mongo.IndexModel{emailIndex, usernameIndex}
}

func GetUserByUsername(username string) (user *types.User, err error) {
	err = collection().FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	return user, err
}

func GetUserByEmail(email string) (user *types.User, err error) {
	err = collection().FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	return user, err
}

// Adds new user into database collection
func AddUser(username, email, passwordHash string) (result *mongo.InsertOneResult, err error) {
	user := &types.User{
		Active:        true,
		Username:      username,
		Email:         email,
		PasswordHash:  passwordHash,
		EmailVerified: false,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	// create indexes for unique constraints
	if _, err = collection().Indexes().CreateMany(context.Background(), getIndexes()); err != nil {
		return nil, err
	}

	return collection().InsertOne(context.Background(), user)
}
