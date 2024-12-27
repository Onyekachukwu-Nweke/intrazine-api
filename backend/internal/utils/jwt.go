package utils

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

var jwtSecret = []byte("missionimpossible")

func GenerateResetToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Minute * 15).Unix(), // 15-minute expiration
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateResetToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", errors.New("invalid user_id in token")
	}

	return userID, nil
}

func ValidateRefreshToken(tokenString string) (string, error) {
	// Similar logic to `ValidateResetToken`
	return ValidateResetToken(tokenString) // Example reuse
}
