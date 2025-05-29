package project

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/validation"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/controller/model/request"
	projectModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/projects"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/view"
)

var (
	ProjectDomainInterface projectModel.ProjectDomainInterface
)

func (pc *projectControllerInterface) CreateProject(c *gin.Context) {
	log.Println("Init CreateProject controller")
	var projectRequest request.ProjectRequest

	if err := c.ShouldBindJSON(&projectRequest); err != nil {
		errRest := validation.ValidateRequestError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := projectModel.NewProjectDomain(
		projectRequest.Name,
		projectRequest.IdAsana,
	)

	domainResult, err := pc.service.CreateProjectServices(domain)
	if err != nil {
		log.Println("Error creating project:", err)
		c.JSON(err.Code, err)
		return
	}

	if domainResult == nil {
		log.Println("Error: domainResult is nil")
		c.JSON(http.StatusInternalServerError, "Project creation failed, domainResult is nil")
		return
	}

	c.JSON(http.StatusOK, view.ConvertProjectDomainToResponse(domainResult))
}
