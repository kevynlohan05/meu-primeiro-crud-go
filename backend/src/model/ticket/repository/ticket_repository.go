package repository

import (
	"database/sql"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	ticketModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket"
)

func NewTicketRepository(database *sql.DB) TicketRepository {
	return &ticketRepository{
		databaseConnection: database,
	}
}

type ticketRepository struct {
	databaseConnection *sql.DB
}

type TicketRepository interface {
	CreateTicket(ticketDomain ticketModel.TicketDomainInterface) (ticketModel.TicketDomainInterface, *rest_err.RestErr)
	UpdateTicket(ticketId string, ticketDomain ticketModel.TicketDomainInterface) *rest_err.RestErr
	UpdateAsanaTaskID(ticketId string, taskID string) *rest_err.RestErr
	UpdateTicketStatus(ticketId string, status string) *rest_err.RestErr
	DeleteTicket(ticketId string) *rest_err.RestErr
	DeleteComment(ticketId string, commentId string) *rest_err.RestErr
	FindAllTicketsByEmail(email string) ([]ticketModel.TicketDomainInterface, *rest_err.RestErr)
	FindTicketById(id string) (ticketModel.TicketDomainInterface, *rest_err.RestErr)
	FindAllTickets() ([]ticketModel.TicketDomainInterface, *rest_err.RestErr)
	FindCommentsByTicketID(ticketId string) ([]ticketModel.CommentDomain, *rest_err.RestErr)
	FindCommentsByEmail(email string) ([]ticketModel.CommentDomain, *rest_err.RestErr)
	AddComment(ticketId string, comment ticketModel.CommentDomain) *rest_err.RestErr
}
