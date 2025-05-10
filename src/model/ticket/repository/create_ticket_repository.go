package repository

import (
	"context"
	"log"
	"os"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model/converter"
	ticketModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (tr *ticketRepository) CreateTicket(ticketDomain ticketModel.TicketDomainInterface) (ticketModel.TicketDomainInterface, *rest_err.RestErr) {
	collection_name := os.Getenv(MONGODB_TICKET_COLLECTION)
	collection := tr.databaseConnection.Collection(collection_name)

	log.Println("Converting domain to entity")

	value := converter.ConvertTicketDomainToEntity(ticketDomain)

	// Log para verificar se o valor está correto
	if value == nil {
		log.Println("Error: Conversion to entity failed")
		return nil, rest_err.NewInternalServerError("Failed to convert domain to entity")
	}

	log.Println("Inserting ticket into MongoDB")
	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		log.Println("Error inserting ticket into MongoDB:", err)
		return nil, rest_err.NewInternalServerError(err.Error())

	}

	log.Println("Ticket inserted successfully into MongoDB")
	value.ID = result.InsertedID.(primitive.ObjectID)

	log.Println("Converting entity back to domain")
	return converter.ConvertTicketEntityToDomain(*value), nil
}
