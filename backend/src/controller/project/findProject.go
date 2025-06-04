package project

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/controller/model/response"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/view"
)

// FindProjectById handles the retrieval of a project by its ID.
func (pc *projectControllerInterface) FindProjectById(c *gin.Context) {
	log.Println("[FindProjectById] Init")

	projectId := c.Param("projectId")

	projectDomain, err := pc.service.FindProjectByIdServices(projectId)
	if err != nil {
		log.Println("[FindProjectById] Service error:", err)
		c.JSON(err.Code, err)
		return
	}

	if projectDomain == nil {
		log.Println("[FindProjectById] Project not found:", projectId)
		c.JSON(http.StatusNotFound, rest_err.NewNotFoundError("Projeto não encontrado"))
		return
	}

	log.Println("[FindProjectById] Project found:", projectId)
	c.JSON(http.StatusOK, view.ConvertProjectDomainToResponse(projectDomain))
}

// FindProjectByName handles the retrieval of a project by its name.
func (pc *projectControllerInterface) FindProjectByName(c *gin.Context) {
	log.Println("[FindProjectByName] Init")

	projectName := c.Param("projectName")

	projectDomain, err := pc.service.FindProjectByNameServices(projectName)
	if err != nil {
		log.Println("[FindProjectByName] Service error:", err)
		c.JSON(err.Code, err)
		return
	}

	if projectDomain == nil {
		log.Println("[FindProjectByName] Project not found:", projectName)
		c.JSON(http.StatusNotFound, rest_err.NewNotFoundError("Projeto não encontrado"))
		return
	}

	log.Println("[FindProjectByName] Project found:", projectName)
	c.JSON(http.StatusOK, view.ConvertProjectDomainToResponse(projectDomain))
}

// FindAllProjects handles the retrieval of all registered projects.
func (pc *projectControllerInterface) FindAllProjects(c *gin.Context) {
	log.Println("[FindAllProjects] Init")

	projects, err := pc.service.FindAllProjectsServices()
	if err != nil {
		log.Println("[FindAllProjects] Service error:", err)
		c.JSON(err.Code, err)
		return
	}

	if len(projects) == 0 {
		log.Println("[FindAllProjects] No projects found")
		c.JSON(http.StatusNotFound, rest_err.NewNotFoundError("Nenhum projeto encontrado"))
		return
	}

	log.Printf("[FindAllProjects] %d projects found\n", len(projects))

	var projectsResponse []response.ProjectResponse
	for _, project := range projects {
		projectsResponse = append(projectsResponse, view.ConvertProjectDomainToResponse(project))
	}

	c.JSON(http.StatusOK, projectsResponse)
}
