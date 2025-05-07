package repository

import (
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model"
	"go.mongodb.org/mongo-driver/mongo"
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
}
