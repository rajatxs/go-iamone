package util

import (
	"os"
	"time"

	"github.com/rajatxs/go-iamone/types"

	"github.com/golang-jwt/jwt/v5"
)

// Returns Access Token Secret Key
func getAccessTokenSecret() []byte {
	return []byte(os.Getenv("JWT_ACCESS_TOKEN_SECRET"))
}

// // Returns Refresh Token Secret Key
// func getRefreshTokenSecret() string {
// 	return os.Getenv("JWT_REFRESH_TOKEN_SECRET")
// }

// Generates JWT Access and Refresh Token by using default signing method
func GenerateAuthToken(name string, id string, admin bool) (*types.AuthToken, error) {
	var (
		gen *jwt.Token
		err error
	)

	now := time.Now()
	claims := jwt.MapClaims{}
	token := &types.AuthToken{}

	claims["iss"] = os.Getenv("IAMONE_SERVER_HOST")
	claims["name"] = name
	claims["sub"] = id
	claims["admin"] = admin
	claims["iat"] = now.Unix()
	claims["exp"] = now.Add(time.Hour * 24).Unix()

	gen = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	if token.AccessToken, err = gen.SignedString(getAccessTokenSecret()); err != nil {
		return nil, err
	}

	return token, err
}
