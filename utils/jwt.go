package utils

import (
	"fmt"
	"os"
	"siakad/api/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// ErrMissingJWTSecret is returned when JWT_SECRET env variable is not set.
var ErrMissingJWTSecret = fmt.Errorf("JWT_SECRET environment variable is not set")

func GenerateToken(user *models.User, expired time.Duration) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", ErrMissingJWTSecret
	}

	claims := jwt.MapClaims{
		"user_id":    user.ID,
		"email":      user.Email,
		"name":       user.Name,
		"role":       user.Role,
		"sekolah_id": user.SekolahID,
		"exp":        expired,
		"iat":        time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return nil, ErrMissingJWTSecret
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verify signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	return token, err
}

// Helper function to extract user info from token
func GetUserFromToken(tokenString string) (uint, string, error) {
	token, err := VerifyToken(tokenString)
	if err != nil {
		return 0, "", err
	}

	if !token.Valid {
		return 0, "", fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, "", fmt.Errorf("invalid claims")
	}

	// Extract user_id
	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, "", fmt.Errorf("invalid user_id")
	}
	userID := uint(userIDFloat)

	// Extract role
	role, ok := claims["role"].(string)
	if !ok {
		return 0, "", fmt.Errorf("invalid role")
	}

	return userID, role, nil
}
