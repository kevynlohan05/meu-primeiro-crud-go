package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/validation"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/controller/model/request"
	userModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/view"
)

func (uc *userControllerInterface) LoginUser(c *gin.Context) {
	log.Println("Init LoginUser controller")
	var userRequest request.UserLoginRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		errRest := validation.ValidateRequestError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := userModel.NewUserLoginDomain(
		userRequest.Email,
		userRequest.Password,
	)

	domainResult, token, err := uc.service.LoginUserServices(domain)
	if err != nil {
		log.Println("Error calling login service:", err)
		c.JSON(err.Code, err)
		return
	}

	c.Header("Authorization", token)
	c.JSON(http.StatusOK, view.ConvertUserDomainToResponse(domainResult))
}
