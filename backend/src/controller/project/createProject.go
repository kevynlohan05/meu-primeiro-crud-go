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

// CreateProject handles the creation of a new project.
// It validates the incoming JSON request, converts it to a domain model,
// passes it to the service layer, and returns the response or error.
func (pc *projectControllerInterface) CreateProject(c *gin.Context) {
	log.Println("[CreateProject] Init CreateProject controller")

	var projectRequest request.ProjectRequest

	// Parse and validate the JSON input
	if err := c.ShouldBindJSON(&projectRequest); err != nil {
		log.Println("[CreateProject] Error binding JSON:", err)
		errRest := validation.ValidateRequestError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	// Convert request to domain model
	domain := projectModel.NewProjectDomain(
		projectRequest.Name,
		projectRequest.IdAsana,
	)

	// Call service layer to create the project
	domainResult, err := pc.service.CreateProjectServices(domain)
	if err != nil {
		log.Println("[CreateProject] Error creating project:", err)
		c.JSON(err.Code, err)
		return
	}

	if domainResult == nil {
		log.Println("[CreateProject] domainResult is nil - unexpected internal error")
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create project"})
		return
	}

	// Successfully created project
	log.Println("[CreateProject] Project created successfully:", domainResult.GetName())
	c.JSON(http.StatusOK, view.ConvertProjectDomainToResponse(domainResult))
}
