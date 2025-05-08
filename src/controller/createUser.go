package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/validation"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/controller/model/request"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/view"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	log.Println("Init CreateUser controller")
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		errRest := validation.ValidateRequestError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Name,
		userRequest.Email,
		userRequest.Password,
		userRequest.Department,
		userRequest.Role,
	)

	domainResult, err := uc.service.CreateUserServices(domain)
	if err != nil {
		log.Println("Error creating user:", err)
		c.JSON(err.Code, err)
		return
	}

	if domainResult == nil {
		log.Println("Error: domainResult is nil")
		c.JSON(http.StatusInternalServerError, "User creation failed, domainResult is nil")
		return
	}

	c.JSON(http.StatusOK, view.ConvertUserDomainToResponse(domainResult))
}
