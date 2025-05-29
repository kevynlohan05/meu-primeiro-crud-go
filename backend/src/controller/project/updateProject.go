package project

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/validation"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/controller/model/request"
	projectModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/projects"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (pc *projectControllerInterface) UpdateProject(c *gin.Context) {
	log.Println("Init UpdateProject controller")
	var projectRequest request.ProjectUpdateRequest

	if err := c.ShouldBindJSON(&projectRequest); err != nil {
		errRest := validation.ValidateRequestError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	projectId := c.Param("projectId")
	if _, err := primitive.ObjectIDFromHex(projectId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid project ID format")
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := projectModel.NewProjectUpdateDomain(
		projectRequest.Name,
		projectRequest.IdAsana,
	)

	err := pc.service.UpdateProject(projectId, domain)
	if err != nil {
		log.Println("Error update user:", err)
		c.JSON(err.Code, err)
		return
	}

	c.Status(http.StatusOK)
}
