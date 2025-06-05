package project

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/controller/model/response"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/view"
)

// FindProjectById handles retrieving a project by its unique ID.
func (pc *projectControllerInterface) FindProjectById(c *gin.Context) {
	log.Println("[FindProjectById] Starting controller")

	projectId := c.Param("projectId")

	projectDomain, err := pc.service.FindProjectByIdServices(projectId)
	if err != nil {
		log.Println("[FindProjectById] Service error:", err)
		c.JSON(err.Code, err)
		return
	}

	if projectDomain == nil {
		log.Println("[FindProjectById] Project not found:", projectId)
		c.JSON(http.StatusNotFound, rest_err.NewNotFoundError("Project not found"))
		return
	}

	log.Println("[FindProjectById] Project retrieved successfully:", projectId)
	c.JSON(http.StatusOK, view.ConvertProjectDomainToResponse(projectDomain))
}

// FindProjectByName handles retrieving a project by its name.
func (pc *projectControllerInterface) FindProjectByName(c *gin.Context) {
	log.Println("[FindProjectByName] Starting controller")

	projectName := c.Param("projectName")

	projectDomain, err := pc.service.FindProjectByNameServices(projectName)
	if err != nil {
		log.Println("[FindProjectByName] Service error:", err)
		c.JSON(err.Code, err)
		return
	}

	if projectDomain == nil {
		log.Println("[FindProjectByName] Project not found:", projectName)
		c.JSON(http.StatusNotFound, rest_err.NewNotFoundError("Project not found"))
		return
	}

	log.Println("[FindProjectByName] Project retrieved successfully:", projectName)
	c.JSON(http.StatusOK, view.ConvertProjectDomainToResponse(projectDomain))
}

// FindAllProjects handles retrieving all registered projects in the system.
func (pc *projectControllerInterface) FindAllProjects(c *gin.Context) {
	log.Println("[FindAllProjects] Starting controller")

	projects, err := pc.service.FindAllProjectsServices()
	if err != nil {
		log.Println("[FindAllProjects] Service error:", err)
		c.JSON(err.Code, err)
		return
	}

	if len(projects) == 0 {
		log.Println("[FindAllProjects] No projects found")
		c.JSON(http.StatusNotFound, rest_err.NewNotFoundError("No projects found"))
		return
	}

	log.Printf("[FindAllProjects] %d projects retrieved successfully\n", len(projects))

	var projectsResponse []response.ProjectResponse
	for _, project := range projects {
		projectsResponse = append(projectsResponse, view.ConvertProjectDomainToResponse(project))
	}

	c.JSON(http.StatusOK, projectsResponse)
}
