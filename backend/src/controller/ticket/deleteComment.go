package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
)

func (tc *ticketControllerInterface) DeleteComment(c *gin.Context) {
	log.Println("Init DeleteComment controller")

	ticketId := c.Param("ticketId")
	commentId := c.Param("commentId")

	if ticketId == "" || commentId == "" {
		errorMessage := rest_err.NewBadRequestError("Ticket ID and Comment ID are required")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userEmail := c.GetString("userEmail")
	if userEmail == "" {
		errorMessage := rest_err.NewUnauthorizedError("User not authenticated")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	err := tc.service.DeleteComment(ticketId, commentId, userEmail)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}
