package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/controller/model/response"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/view"
)

func (tc *ticketControllerInterface) FindTicketById(c *gin.Context) {
	ticketId := c.Param("ticketId")

	ticketDomain, err := tc.service.FindTicketByIdServices(ticketId)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	if ticketDomain == nil {
		errorMessage := rest_err.NewNotFoundError("Ticket not found")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	c.JSON(http.StatusOK, view.ConvertTicketDomainToResponse(ticketDomain))
}

func (tc *ticketControllerInterface) FindAllTicketsByEmail(c *gin.Context) {
	ticketEmail := c.Param("ticketEmail")

	// Validação simples de e-mail (opcional)
	if ticketEmail == "" {
		errorMessage := rest_err.NewBadRequestError("Email do solicitante é obrigatório")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	tickets, err := tc.service.FindAllTicketsByEmail(ticketEmail)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	if len(tickets) == 0 {
		errorMessage := rest_err.NewNotFoundError("No tickets found for the provided email")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

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

	if len(tickets) == 0 {
		errorMessage := rest_err.NewNotFoundError("No tickets found")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	var ticketResponses []response.TicketResponse
	for _, ticket := range tickets {
		ticketResponses = append(ticketResponses, view.ConvertTicketDomainToResponse(ticket))
	}

	c.JSON(http.StatusOK, ticketResponses)
}

func (tc *ticketControllerInterface) FindAllTicketsByEmailAndStatus(c *gin.Context) {
	ticketEmail := c.Param("ticketEmail")
	ticketStatus := c.Param("ticketStatus")

	if ticketEmail == "" {
		errorMessage := rest_err.NewBadRequestError("Email do solicitante é obrigatório")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	if ticketStatus == "" {
		errorMessage := rest_err.NewBadRequestError("Status é obrigatório")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	tickets, err := tc.service.FindAllTicketsByEmailAndStatus(ticketEmail, ticketStatus)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	if len(tickets) == 0 {
		errorMessage := rest_err.NewNotFoundError("No tickets found for the provided email and status")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	var ticketResponses []response.TicketResponse
	for _, ticket := range tickets {
		ticketResponses = append(ticketResponses, view.ConvertTicketDomainToResponse(ticket))
	}

	c.JSON(http.StatusOK, ticketResponses)
}
