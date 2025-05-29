package model

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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
		"phone":      ud.phone,
		"enterprise": ud.enterprise,
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

func VerifyTokenMiddleware(c *gin.Context) {
	secret := os.Getenv(JWT_SECRET_KEY)
	tokenValue := RemoveBearerPrefix(c.Request.Header.Get("Authorization"))

	token, err := jwt.Parse(tokenValue, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}
		c.Abort()
		return nil, rest_err.NewBadRequestError("invalid token")
	})

	if err != nil {
		errRest := rest_err.NewUnauthorizedError("invalid token")
		c.JSON(errRest.Code, errRest)
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		errRest := rest_err.NewUnauthorizedError("invalid token")
		c.JSON(errRest.Code, errRest)
		c.Abort()
		return
	}

	userID := claims["id"].(string)
	userName := claims["name"].(string)
	userEmail := claims["email"].(string)
	userDepartment := claims["department"].(string)
	userRole := claims["role"].(string)

	// Salvando no contexto para uso posterior

	c.Set("userID", userID)
	c.Set("userName", userName)
	c.Set("userEmail", userEmail)
	c.Set("userDepartment", userDepartment)
	c.Set("role", userRole)

	c.Next()
}

func AdminOnlyMiddleware(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists || role != "admin" {
		errRest := rest_err.NewUnauthorizedError("Acesso permitido apenas para administradores")
		c.JSON(errRest.Code, errRest)
		c.Abort()
		return
	}
	c.Next()
}
