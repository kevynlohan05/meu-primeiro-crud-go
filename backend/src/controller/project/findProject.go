package project

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/controller/model/response"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/view"
)

func (pc *projectControllerInterface) FindProjectById(c *gin.Context) {
	projectId := c.Param("projectId")

	projectDomain, err := pc.service.FindProjectByIdServices(projectId)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	if projectDomain == nil {
		log.Println("Project not found")
		c.JSON(http.StatusNotFound, rest_err.NewNotFoundError("Projeto não encontrado"))
		return
	}

	c.JSON(http.StatusOK, view.ConvertProjectDomainToResponse(projectDomain))
}

func (pc *projectControllerInterface) FindProjectByName(c *gin.Context) {
	projectName := c.Param("projectName")

	projectDomain, err := pc.service.FindProjectByNameServices(projectName)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	if projectDomain == nil {
		log.Println("Project not found")
		c.JSON(http.StatusNotFound, rest_err.NewNotFoundError("Projeto não encontrado"))
		return
	}

	c.JSON(http.StatusOK, view.ConvertProjectDomainToResponse(projectDomain))

}

func (pc *projectControllerInterface) FindAllProjects(c *gin.Context) {
	projects, err := pc.service.FindAllProjectsServices()
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	if len(projects) == 0 {
		log.Println("No projects found")
		c.JSON(http.StatusNotFound, rest_err.NewNotFoundError("Nenhum projeto encontrado"))
		return
	}

	var projectsResponse []response.ProjectResponse
	for _, project := range projects {
		projectsResponse = append(projectsResponse, view.ConvertProjectDomainToResponse(project))
	}

	c.JSON(http.StatusOK, projectsResponse)
}
