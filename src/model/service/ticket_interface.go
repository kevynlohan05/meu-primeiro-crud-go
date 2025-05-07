package service

import (
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model/repository"
)

func NewTicketDomainService(ticketRepository repository.TicketRepository) TicketDomainService {
	return &ticketDomainService{ticketRepository}
}

type ticketDomainService struct {
	ticketRepository repository.TicketRepository
}

type TicketDomainService interface {
	CreateTicket(model.TicketDomainInterface) (model.TicketDomainInterface, *rest_err.RestErr)
	UpdateTicket(string, model.TicketDomainInterface) *rest_err.RestErr
	DeleteTicket(string) *rest_err.RestErr
	FindTicketById(string) (*model.TicketDomainInterface, *rest_err.RestErr)
}
