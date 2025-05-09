package controller

import (
	"log"
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/view"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc *userControllerInterface) FindUserById(c *gin.Context) {
	userId := c.Param("userId")

	_, err := model.VerifyToken(c.Request.Header.Get("Authorization"))
	if err != nil {
		log.Println("Error trying to verify token:", err)
		c.JSON(err.Code, err)
		return
	}

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errorMessage := rest_err.NewBadRequestError("Invalid user ID format")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByIdServices(userId)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertUserDomainToResponse(userDomain))
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	userEmail := c.Param("userEmail")

	_, err := model.VerifyToken(c.Request.Header.Get("Authorization"))
	if err != nil {
		log.Println("Error trying to verify token:", err)
		c.JSON(err.Code, err)
		return
	}

	if _, err := mail.ParseAddress(userEmail); err != nil {
		errorMessage := rest_err.NewBadRequestError("Invalid user Email format")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmailServices(userEmail)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertUserDomainToResponse(userDomain))

}
