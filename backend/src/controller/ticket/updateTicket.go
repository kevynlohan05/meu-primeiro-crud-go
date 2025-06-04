package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/validation"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/controller/model/request"
	ticketModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (tc *ticketControllerInterface) UpdateTicket(c *gin.Context) {
	log.Println("Init UpdateTicket controller")

	var ticketRequest request.TicketUpdateRequest
	if err := c.ShouldBindJSON(&ticketRequest); err != nil {
		log.Printf("Validation error in UpdateTicket: %v\n", err)
		errRest := validation.ValidateRequestError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	ticketId := c.Param("ticketId")
	log.Printf("Updating ticket with ID: %s\n", ticketId)

	if _, err := primitive.ObjectIDFromHex(ticketId); err != nil {
		log.Printf("Invalid ticket ID format: %s\n", ticketId)
		errRest := rest_err.NewBadRequestError("Invalid ticket ID format")
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := ticketModel.NewTicketUpdateDomain(
		ticketRequest.Title,
		ticketRequest.Description,
		ticketRequest.RequestType,
		ticketRequest.Priority,
		ticketRequest.Status,
	)

	err := tc.service.UpdateTicket(ticketId, domain)
	if err != nil {
		log.Printf("Error updating ticket %s: %v\n", ticketId, err)
		c.JSON(err.Code, err)
		return
	}

	log.Printf("Ticket %s updated successfully\n", ticketId)
	c.JSON(http.StatusOK, gin.H{"message": "Ticket updated successfully"})
}
