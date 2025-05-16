package repository

import (
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	ticketModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGODB_TICKET_COLLECTION = "MONGODB_TICKET_COLLECTION"
)

func NewTicketRepository(database *mongo.Database) TicketRepository {
	return &ticketRepository{
		database,
	}

}

type ticketRepository struct {
	databaseConnection *mongo.Database
}

type TicketRepository interface {
	CreateTicket(ticketDomain ticketModel.TicketDomainInterface) (ticketModel.TicketDomainInterface, *rest_err.RestErr)

	UpdateTicket(ticketId string, ticketDomain ticketModel.TicketDomainInterface) *rest_err.RestErr

	UpdateAsanaTaskID(ticketId string, taskID string) *rest_err.RestErr

	DeleteTicket(ticketId string) *rest_err.RestErr

	FindAllTicketsByEmail(email string) ([]ticketModel.TicketDomainInterface, *rest_err.RestErr)

	FindTicketById(id string) (ticketModel.TicketDomainInterface, *rest_err.RestErr)

	FindAllTickets() ([]ticketModel.TicketDomainInterface, *rest_err.RestErr)

	AddComment(ticketId string, comment ticketModel.CommentDomain) *rest_err.RestErr
}
