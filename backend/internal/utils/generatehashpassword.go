package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword generates a bcrypt hash of the password
func HashPassword(password string) (string, error) {
	hash_pass, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(hash_pass), err
}

// CheckPasswordHash compares a bcrypt hashed password with its possible plaintext equivalent
func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateJWT generates a JWT token for authenticated users
// Implementation of JWT Generation