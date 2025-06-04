package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
)

func (tc *ticketControllerInterface) DeleteTicket(c *gin.Context) {
	log.Println("Init DeleteTicket controller")

	ticketId := c.Param("ticketId")
	if ticketId == "" {
		log.Println("Ticket ID is required")
		c.JSON(http.StatusBadRequest, rest_err.NewBadRequestError("Ticket ID is required"))
		return
	}

	err := tc.service.DeleteTicket(ticketId)
	if err != nil {
		log.Println("Error update ticket:", err)
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}
