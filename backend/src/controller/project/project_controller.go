package project

import (
	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model/projects/service"
)

// NewProjectControllerInterface returns a new instance of the project controller
func NewProjectControllerInterface(serviceInterface service.ProjectDomainService) ProjectControllerInterface {
	return &projectControllerInterface{
		service: serviceInterface,
	}
}

// ProjectControllerInterface defines the methods that a project controller should implement
type ProjectControllerInterface interface {
	CreateProject(c *gin.Context)
	UpdateProject(c *gin.Context)
	DeleteProject(c *gin.Context)
	FindProjectById(c *gin.Context)
	FindProjectByName(c *gin.Context)
	FindAllProjects(c *gin.Context)
}

// projectControllerInterface is the concrete implementation of ProjectControllerInterface
type projectControllerInterface struct {
	service service.ProjectDomainService
}
