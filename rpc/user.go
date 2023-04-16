package rpc

import (
	"errors"
	"log"
	"net/http"

	"github.com/rajatxs/go-iamone/db"
	"github.com/rajatxs/go-iamone/types"
	"github.com/rajatxs/go-iamone/util"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct{}

type UserGenerateTokenArgs struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Generates auth token by given username or email
func (u *User) GenerateToken(r *http.Request, args *UserGenerateTokenArgs, reply *types.AuthToken) (err error) {
	var (
		user  *types.User
		token *types.AuthToken
	)

	// username should have first priority
	if len(args.Username) > 0 {
		user, err = db.GetUserByUsername(args.Username)
	} else if len(args.Email) > 0 {
		user, err = db.GetUserByEmail(args.Email)
	} else {
		return errors.New("require username or email")
	}

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("account not found")
		} else {
			log.Println(err)
			return errors.New("couldn't get your account")
		}
	}

	if !util.ComparePasswordHash(user.PasswordHash, args.Password) {
		return errors.New("incorrect password")
	}

	if token, err = util.GenerateAuthToken(user.Username, user.ID.Hex(), false); err != nil {
		log.Println(err)
		return errors.New("couldn't generate identity")
	}

	*reply = *token
	return nil
}

// Register new user from username, email and new password
// Returns generated auth token after registration
func (u *User) Register(r *http.Request, args *UserGenerateTokenArgs, reply *types.AuthToken) (err error) {
	var (
		passwordHash string
		userId       string
		result       *mongo.InsertOneResult
		token        *types.AuthToken
	)

	if passwordHash, err = util.GeneratePasswordHash(args.Password); err != nil {
		log.Println(err)
		return errors.New("couldn't encode given password")
	}

	if result, err = db.AddUser(args.Username, args.Email, passwordHash); err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return errors.New("username or email is already in use")
		} else {
			log.Println(err)
			return errors.New("couldn't register account")
		}
	}

	if result.InsertedID == nil {
		return errors.New("something went wrong")
	}

	userId = result.InsertedID.(primitive.ObjectID).Hex()

	if token, err = util.GenerateAuthToken(args.Username, userId, false); err != nil {
		log.Println(err)
		return errors.New("couldn't generate identity")
	}

	*reply = *token
	return nil
}
