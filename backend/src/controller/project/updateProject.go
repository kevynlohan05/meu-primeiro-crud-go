package project

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/validation"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/controller/model/request"
	projectModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/projects"
)

// UpdateProject handles updating the name and Asana ID of a project.
func (pc *projectControllerInterface) UpdateProject(c *gin.Context) {
	log.Println("[UpdateProject] Starting controller")

	var projectRequest request.ProjectUpdateRequest

	// Validate incoming JSON body
	if err := c.ShouldBindJSON(&projectRequest); err != nil {
		log.Println("[UpdateProject] Error binding JSON:", err)
		errRest := validation.ValidateRequestError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	projectId := c.Param("projectId")

	// Create domain object for update
	domain := projectModel.NewProjectUpdateDomain(
		projectRequest.Name,
		projectRequest.IdAsana,
	)

	// Execute service update
	err := pc.service.UpdateProject(projectId, domain)
	if err != nil {
		log.Println("[UpdateProject] Error updating project:", err)
		c.JSON(err.Code, err)
		return
	}

	log.Println("[UpdateProject] Project updated successfully:", projectId)
	c.Status(http.StatusOK)
}
