package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
)

func (tc *ticketControllerInterface) DeleteComment(c *gin.Context) {
	log.Println("Start DeleteComment controller")

	ticketId := c.Param("ticketId")
	commentId := c.Param("commentId")

	if ticketId == "" || commentId == "" {
		errRest := rest_err.NewBadRequestError("Ticket ID and Comment ID are required")
		c.JSON(errRest.Code, errRest)
		return
	}

	userEmail := c.GetString("userEmail")
	if userEmail == "" {
		errRest := rest_err.NewUnauthorizedError("User not authenticated")
		c.JSON(errRest.Code, errRest)
		return
	}

	if err := tc.service.DeleteComment(ticketId, commentId, userEmail); err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}
