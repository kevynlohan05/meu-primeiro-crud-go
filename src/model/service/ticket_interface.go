package service

import (
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model"
)

func NewTicketDomainService() TicketDomainService {
	return &ticketDomainService{}
}

type ticketDomainService struct {
}

type TicketDomainService interface {
	CreateTicket(model.TicketDomainInterface) *rest_err.RestErr
	UpdateTicket(string, model.TicketDomainInterface) *rest_err.RestErr
	DeleteTicket(string) *rest_err.RestErr
	FindTicketById(string) (*model.TicketDomainInterface, *rest_err.RestErr)
}
