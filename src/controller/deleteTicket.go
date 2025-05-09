package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (tc *ticketControllerInterface) DeleteTicket(c *gin.Context) {
	log.Println("Init DeleteTicket controller")

	ticketId := c.Param("ticketId")
	if _, err := primitive.ObjectIDFromHex(ticketId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid user ID format")
		c.JSON(errRest.Code, errRest)
		return
	}

	err := tc.service.DeleteTicket(ticketId)
	if err != nil {
		log.Println("Error update ticket:", err)
		c.JSON(err.Code, err)
		return
	}

	c.Status(http.StatusOK)
}
