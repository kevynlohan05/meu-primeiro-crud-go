package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket/service"
)

func NewTicketControllerInterface(serviceInterface service.TicketDomainService) TicketControllerInterface {
	return &ticketControllerInterface{
		service: serviceInterface,
	}
}

type TicketControllerInterface interface {
	CreateTicket(c *gin.Context)
	UpdateTicket(c *gin.Context)
	DeleteTicket(c *gin.Context)
	FindTicketById(c *gin.Context)
	FindTicketByEmail(c *gin.Context)
}

type ticketControllerInterface struct {
	service service.TicketDomainService
}
