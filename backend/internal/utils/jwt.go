package utils

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

// contextKey is an unexported type for context keys defined in this package.
// This prevents collisions with keys defined in other packages.
type contextKey string

// userContextKey is the key for user ID values in the context.
const userContextKey contextKey = "user_id"

func JWTAuth(original func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header["Authorization"]
		if authHeader == nil {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}

		// Bearer: token-string
		authHeaderParts := strings.Split(authHeader[0], " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}

		userID, err := validateToken(authHeaderParts[1])
		fmt.Println(userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// Set user_id in the context
		ctx := context.WithValue(r.Context(), userContextKey, userID)
		r = r.WithContext(ctx)

		// Call the original handler
		original(w, r)
	}
}

func validateToken(accessToken string) (string, error) {
	var mySigningKey = []byte("missionimpossible")
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("could not validate auth token")
		}

		return mySigningKey, nil
	})

	if err != nil {
		return "", errors.New("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)
		userID, ok := claims["user_id"].(string)
		fmt.Println(userID)
		if !ok {
			return "", errors.New("invalid token claims")
		}
		return userID, nil
	}

	return "", errors.New("invalid token")
}
