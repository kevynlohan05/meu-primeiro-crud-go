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
		errRest := validation.ValidateRequestError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	ticketId := c.Param("ticketId")
	if _, err := primitive.ObjectIDFromHex(ticketId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid user ID format")
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := ticketModel.NewTicketUpdateDomain(
		ticketRequest.Title,
		ticketRequest.Description,
		ticketRequest.RequestType,
		ticketRequest.Priority,
		ticketRequest.AttachmentURL,
	)

	err := tc.service.UpdateTicket(ticketId, domain)
	if err != nil {
		log.Println("Error update ticket:", err)
		c.JSON(err.Code, err)
		return
	}

	c.Status(http.StatusOK)
}
