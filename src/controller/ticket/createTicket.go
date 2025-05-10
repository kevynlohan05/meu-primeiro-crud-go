package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/validation"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/controller/model/request"
	ticketModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/view"
)

var (
	TicketDomainInterface ticketModel.TicketDomainInterface
)

func (tc *ticketControllerInterface) CreateTicket(c *gin.Context) {

	log.Println("Init createTicket controller")
	var ticketRequest request.TicketRequest

	if err := c.ShouldBindJSON(&ticketRequest); err != nil {
		errRest := validation.ValidateRequestError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := ticketModel.NewTicketDomain(
		ticketRequest.Title,
		ticketRequest.Description,
		ticketRequest.RequestType,
		ticketRequest.Priority,
		ticketRequest.AttachmentURL,
	)

	domainResult, err := tc.service.CreateTicket(domain)
	if err != nil {
		log.Println("Error creating ticket:", err)
		c.JSON(err.Code, err)
		return
	}

	if domainResult == nil {
		log.Println("Error: domainResult is nil")
		c.JSON(http.StatusInternalServerError, "Ticket creation failed, domainResult is nil")
		return
	}

	c.JSON(http.StatusOK, view.ConvertTicketDomainToResponse(domainResult))
}
