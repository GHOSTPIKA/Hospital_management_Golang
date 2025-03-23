package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWTSecretKey - Change this for production!
var JWTSecretKey = []byte("your-secret-key")

// Claims struct for JWT payload
type Claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

// GenerateJWT creates a new JWT token
func GenerateJWT(userID, role string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // 1-day expiration
	claims := &Claims{
		UserID: userID,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTSecretKey)
}

// ValidateJWT checks if a JWT token is valid
func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JWTSecretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}

	return claims, nil
}
