package util

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func ExtractSubFromToken(tokenString string) (string, error) {
	parser := jwt.NewParser()
	claims := jwt.MapClaims{}
	_, _, err := parser.ParseUnverified(tokenString, claims)
	if err != nil {
		return "", err
	}
	sub, ok := claims["sub"]
	if !ok {
		return "", fmt.Errorf("sub not found in token")
	}
	return fmt.Sprintf("%v", sub), nil
}
