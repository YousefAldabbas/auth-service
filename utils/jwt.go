package utils

import (
	"time"

	"github.com/YousefAldabbas/auth-service/pkg/model"
	"github.com/golang-jwt/jwt"
)

// GenerateJWT generates a JWT token with the provided UUID in the claims
func GenerateJWT(user model.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uuid"] = user.UUID
	claims["username"] = user.Username
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 3).Unix()

	tokenString, err := token.SignedString([]byte("STATIC_FOR_NOW"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
