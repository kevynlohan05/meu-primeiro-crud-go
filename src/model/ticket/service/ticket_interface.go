package service

import (
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	ticketModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket"
	repositoryTicket "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket/repository"
)

func NewTicketDomainService(ticketRepository repositoryTicket.TicketRepository) TicketDomainService {
	return &ticketDomainService{ticketRepository}
}

type ticketDomainService struct {
	ticketRepository repositoryTicket.TicketRepository
}

type TicketDomainService interface {
	CreateTicket(ticketModel.TicketDomainInterface) (ticketModel.TicketDomainInterface, *rest_err.RestErr)
	UpdateTicket(string, ticketModel.TicketDomainInterface) *rest_err.RestErr
	DeleteTicket(string) *rest_err.RestErr
	FindTicketByIdServices(id string) (ticketModel.TicketDomainInterface, *rest_err.RestErr)
	UpdateAsanaTaskID(ticketId string, taskID string) *rest_err.RestErr
	FindAllTicketsByUser(email string) ([]ticketModel.TicketDomainInterface, *rest_err.RestErr)
	FindAllTickets() ([]ticketModel.TicketDomainInterface, *rest_err.RestErr)
	AddComment(ticketId string, comment ticketModel.CommentDomain) *rest_err.RestErr
}
