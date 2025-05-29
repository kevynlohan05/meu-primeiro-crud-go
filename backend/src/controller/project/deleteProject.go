package project

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
)

func (pc *projectControllerInterface) DeleteProject(c *gin.Context) {
	log.Println("Init DeleteProject controller")

	projectId := c.Param("projectId")

	if projectId == "" {
		log.Println("projectId not provided")
		c.JSON(http.StatusBadRequest, rest_err.NewBadRequestError("Parâmetro projectId é obrigatório"))
		return
	}

	err := pc.service.DeleteProject(projectId)
	if err != nil {
		log.Println("Error delete project:", err)
		c.JSON(err.Code, err)
		return
	}

	c.Status(http.StatusOK)
}
