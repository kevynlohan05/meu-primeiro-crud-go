package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/validation"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/controller/model/request"
	userModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	log.Println("Init UpdateUser controller")
	var userRequest request.UserUpdateRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		errRest := validation.ValidateRequestError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid user ID format")
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := userModel.NewUserUpdateDomain(
		userRequest.Name,
		userRequest.Department,
	)

	err := uc.service.UpdateUser(userId, domain)
	if err != nil {
		log.Println("Error update user:", err)
		c.JSON(err.Code, err)
		return
	}

	c.Status(http.StatusOK)
}
