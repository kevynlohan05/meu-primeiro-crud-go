package project

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
)

// DeleteProject handles the deletion of a project by its ID.
// It validates the input parameter, calls the service layer,
// and returns an appropriate HTTP response.
func (pc *projectControllerInterface) DeleteProject(c *gin.Context) {
	log.Println("[DeleteProject] Init DeleteProject controller")

	projectId := c.Param("projectId")

	// Validate required parameter
	if projectId == "" {
		log.Println("[DeleteProject] projectId not provided")
		c.JSON(http.StatusBadRequest, rest_err.NewBadRequestError("Parâmetro projectId é obrigatório"))
		return
	}

	// Call the service to delete the project
	err := pc.service.DeleteProject(projectId)
	if err != nil {
		log.Println("[DeleteProject] Error deleting project:", err)
		c.JSON(err.Code, err)
		return
	}

	log.Println("[DeleteProject] Project deleted successfully:", projectId)
	c.JSON(http.StatusOK, gin.H{"message": "Projeto deletado com sucesso"})
}
