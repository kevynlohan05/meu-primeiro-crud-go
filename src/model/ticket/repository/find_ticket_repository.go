package repository

import (
	"context"
	"os"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model/converter"
	ticketModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket"
	ticketEntity "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket/repository/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (tr *ticketRepository) FindTicketById(id string) (ticketModel.TicketDomainInterface, *rest_err.RestErr) {
	collection_name := os.Getenv(MONGODB_TICKET_COLLECTION)
	collection := tr.databaseConnection.Collection(collection_name)

	ticketEntity := &ticketEntity.TicketEntity{}

	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectId}}
	err := collection.FindOne(context.Background(), filter).Decode(ticketEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := "Ticket not found"
			return nil, rest_err.NewNotFoundError(errorMessage)
		}

		errorMessage := "Error while trying to find ticket: " + err.Error()
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	return converter.ConvertTicketEntityToDomain(*ticketEntity), nil
}

func (tr *ticketRepository) FindTicketByEmail(email string) (ticketModel.TicketDomainInterface, *rest_err.RestErr) {
	collection_name := os.Getenv(MONGODB_TICKET_COLLECTION)
	collection := tr.databaseConnection.Collection(collection_name)

	ticketEntity := &ticketEntity.TicketEntity{}

	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(context.Background(), filter).Decode(ticketEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := "Ticket not found"
			return nil, rest_err.NewNotFoundError(errorMessage)
		}

		errorMessage := "Error while trying to find ticket: " + err.Error()
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	return converter.ConvertTicketEntityToDomain(*ticketEntity), nil
}
