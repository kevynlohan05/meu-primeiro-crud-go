package model

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
)

var (
	JWT_SECRET_KEY = os.Getenv("JWT_SECRET_KEY")
)

func (ud *userDomain) GenerateToken() (string, *rest_err.RestErr) {

	secret := os.Getenv(JWT_SECRET_KEY)

	claims := jwt.MapClaims{
		"id":         ud.iD,
		"name":       ud.name,
		"email":      ud.email,
		"department": ud.department,
		"role":       ud.role,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", rest_err.NewInternalServerError(
			fmt.Sprintf("error trying to generate jwt token, err=%s", err.Error()))
	}

	return tokenString, nil
}

func VerifyToken(tokenValue string) (UserDomainInterface, *rest_err.RestErr) {
	secret := os.Getenv(JWT_SECRET_KEY)

	token, err := jwt.Parse(RemoveBearerPrefix(tokenValue), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}
		return nil, rest_err.NewBadRequestError("invalid token")
	})

	if err != nil {
		return nil, rest_err.NewUnauthorizedError("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, rest_err.NewUnauthorizedError("invalid token")
	}

	return &userDomain{
		iD:         claims["id"].(string),
		name:       claims["name"].(string),
		email:      claims["email"].(string),
		department: claims["department"].(string),
		role:       claims["role"].(string),
	}, nil
}

func RemoveBearerPrefix(token string) string {
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix(token, "Bearer ")

	}
	return token
}
