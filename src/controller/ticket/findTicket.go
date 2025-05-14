package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/controller/model/response"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/view"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (tc *ticketControllerInterface) FindTicketById(c *gin.Context) {
	ticketId := c.Param("ticketId")

	if _, err := primitive.ObjectIDFromHex(ticketId); err != nil {
		errorMessage := rest_err.NewBadRequestError("Invalid ticket ID format")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	ticketDomain, err := tc.service.FindTicketByIdServices(ticketId)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertTicketDomainToResponse(ticketDomain))
}

func (tc *ticketControllerInterface) FindAllTicketsByUser(c *gin.Context) {
	ticketEmail := c.Param("ticketEmail")

	// Validação simples de e-mail (opcional)
	if ticketEmail == "" {
		errorMessage := rest_err.NewBadRequestError("Email do solicitante é obrigatório")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	tickets, err := tc.service.FindAllTicketsByUser(ticketEmail)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	// Converter slice de domains para slice de responses
	var ticketResponses []response.TicketResponse
	for _, ticket := range tickets {
		ticketResponses = append(ticketResponses, view.ConvertTicketDomainToResponse(ticket))
	}

	c.JSON(http.StatusOK, ticketResponses)
}

func (tc *ticketControllerInterface) FindAllTickets(c *gin.Context) {
	tickets, err := tc.service.FindAllTickets()
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	var ticketResponses []response.TicketResponse
	for _, ticket := range tickets {
		ticketResponses = append(ticketResponses, view.ConvertTicketDomainToResponse(ticket))
	}

	c.JSON(http.StatusOK, ticketResponses)
}
