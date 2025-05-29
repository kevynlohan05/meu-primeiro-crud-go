package project

import (
	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model/projects/service"
)

func NewProjectControllerInterface(serviceInterface service.ProjectDomainService) ProjectControllerInterface {
	return &projectControllerInterface{
		service: serviceInterface,
	}
}

type ProjectControllerInterface interface {
	CreateProject(c *gin.Context)
	UpdateProject(c *gin.Context)
	DeleteProject(c *gin.Context)
	FindProjectById(c *gin.Context)
	FindProjectByName(c *gin.Context)
}

type projectControllerInterface struct {
	service service.ProjectDomainService
}
