package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/validation"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/controller/model/request"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model/service"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/view"
)

var (
	TicketDomainInterface model.TicketDomainInterface
)

func (tc *ticketControllerInterface) CreateTicket(c *gin.Context) {

	log.Println("Init createTicket controller")
	var ticketRequest request.TicketRequest

	if err := c.ShouldBindJSON(&ticketRequest); err != nil {
		errRest := validation.ValidateRequestError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewTicketDomain(
		ticketRequest.Title,
		ticketRequest.Description,
		ticketRequest.RequestType,
		ticketRequest.Priority,
		ticketRequest.AttachmentURL,
	)

	service := service.NewTicketDomainService()

	if err := service.CreateTicket(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertTicketDomainToResponse(domain))
}
