package service

import (
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	projectService "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/projects/service"
	ticketModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket"
	repositoryTicket "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket/repository"
	userService "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user/service"
)

func NewTicketDomainService(userService userService.UserDomainService, ticketRepository repositoryTicket.TicketRepository, projectService projectService.ProjectDomainService) TicketDomainService {
	return &ticketDomainService{
		ticketRepository: ticketRepository,
		userService:      userService,
		projectService:   projectService,
	}
}

type ticketDomainService struct {
	ticketRepository repositoryTicket.TicketRepository
	userService      userService.UserDomainService
	projectService   projectService.ProjectDomainService
}

type TicketDomainService interface {
	CreateTicket(ticketModel.TicketDomainInterface) (ticketModel.TicketDomainInterface, *rest_err.RestErr)
	UpdateTicket(string, ticketModel.TicketDomainInterface) *rest_err.RestErr
	DeleteTicket(string) *rest_err.RestErr
	FindTicketByIdServices(id string) (ticketModel.TicketDomainInterface, *rest_err.RestErr)
	UpdateAsanaTaskID(ticketId string, taskID string) *rest_err.RestErr
	FindAllTicketsByEmail(email string) ([]ticketModel.TicketDomainInterface, *rest_err.RestErr)
	FindAllTickets() ([]ticketModel.TicketDomainInterface, *rest_err.RestErr)
	FindAllTicketsByEmailAndStatus(email string, status string) ([]ticketModel.TicketDomainInterface, *rest_err.RestErr)
	AddComment(ticketId string, comment ticketModel.CommentDomain) *rest_err.RestErr
}
