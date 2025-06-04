package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/controller/model/response"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/view"
)

func (tc *ticketControllerInterface) FindTicketById(c *gin.Context) {
	ticketId := c.Param("ticketId")
	log.Printf("FindTicketById called with ticketId: %s\n", ticketId)

	ticketDomain, err := tc.service.FindTicketByIdServices(ticketId)
	if err != nil {
		log.Printf("Error finding ticket by ID: %v\n", err)
		c.JSON(err.Code, err)
		return
	}

	if ticketDomain == nil {
		log.Printf("Ticket not found: %s\n", ticketId)
		errorMessage := rest_err.NewNotFoundError("Ticket not found")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	log.Printf("Ticket found: %s\n", ticketId)
	c.JSON(http.StatusOK, view.ConvertTicketDomainToResponse(ticketDomain))
}

func (tc *ticketControllerInterface) FindAllTicketsByEmail(c *gin.Context) {
	ticketEmail := c.Param("ticketEmail")
	log.Printf("FindAllTicketsByEmail called with email: %s\n", ticketEmail)

	if ticketEmail == "" {
		log.Println("Requester email is required")
		errorMessage := rest_err.NewBadRequestError("Requester email is required")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	tickets, err := tc.service.FindAllTicketsByEmail(ticketEmail)
	if err != nil {
		log.Printf("Error finding tickets by email: %v\n", err)
		c.JSON(err.Code, err)
		return
	}

	if len(tickets) == 0 {
		log.Printf("No tickets found for email: %s\n", ticketEmail)
		errorMessage := rest_err.NewNotFoundError("No tickets found for the provided email")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	log.Printf("Found %d tickets for email: %s\n", len(tickets), ticketEmail)
	var ticketResponses []response.TicketResponse
	for _, ticket := range tickets {
		ticketResponses = append(ticketResponses, view.ConvertTicketDomainToResponse(ticket))
	}

	c.JSON(http.StatusOK, ticketResponses)
}

func (tc *ticketControllerInterface) FindAllTickets(c *gin.Context) {
	log.Println("FindAllTickets called")

	tickets, err := tc.service.FindAllTickets()
	if err != nil {
		log.Printf("Error finding all tickets: %v\n", err)
		c.JSON(err.Code, err)
		return
	}

	if len(tickets) == 0 {
		log.Println("No tickets found")
		errorMessage := rest_err.NewNotFoundError("No tickets found")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	log.Printf("Found %d tickets\n", len(tickets))
	var ticketResponses []response.TicketResponse
	for _, ticket := range tickets {
		ticketResponses = append(ticketResponses, view.ConvertTicketDomainToResponse(ticket))
	}

	c.JSON(http.StatusOK, ticketResponses)
}

func (tc *ticketControllerInterface) FindAllTicketsByEmailAndStatus(c *gin.Context) {
	ticketEmail := c.Param("ticketEmail")
	ticketStatus := c.Param("ticketStatus")
	log.Printf("FindAllTicketsByEmailAndStatus called with email: %s and status: %s\n", ticketEmail, ticketStatus)

	if ticketEmail == "" {
		log.Println("Requester email is required")
		errorMessage := rest_err.NewBadRequestError("Requester email is required")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	if ticketStatus == "" {
		log.Println("Status is required")
		errorMessage := rest_err.NewBadRequestError("Status is required")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	tickets, err := tc.service.FindAllTicketsByEmailAndStatus(ticketEmail, ticketStatus)
	if err != nil {
		log.Printf("Error finding tickets by email and status: %v\n", err)
		c.JSON(err.Code, err)
		return
	}

	if len(tickets) == 0 {
		log.Printf("No tickets found for email: %s and status: %s\n", ticketEmail, ticketStatus)
		errorMessage := rest_err.NewNotFoundError("No tickets found for the provided email and status")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	log.Printf("Found %d tickets for email: %s and status: %s\n", len(tickets), ticketEmail, ticketStatus)
	var ticketResponses []response.TicketResponse
	for _, ticket := range tickets {
		ticketResponses = append(ticketResponses, view.ConvertTicketDomainToResponse(ticket))
	}

	c.JSON(http.StatusOK, ticketResponses)
}
