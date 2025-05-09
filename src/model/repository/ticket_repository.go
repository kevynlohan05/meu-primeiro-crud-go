package repository

import (
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model"
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
	CreateTicket(ticketDomain model.TicketDomainInterface) (model.TicketDomainInterface, *rest_err.RestErr)

	UpdateTicket(ticketId string, ticketDomain model.TicketDomainInterface) *rest_err.RestErr

	DeleteTicket(ticketId string) *rest_err.RestErr

	FindTicketByEmail(email string) (model.TicketDomainInterface, *rest_err.RestErr)

	FindTicketById(id string) (model.TicketDomainInterface, *rest_err.RestErr)
}
