package repository

import (
	"context"
	"os"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model/repository/entity"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (tr *ticketRepository) FindTicketByEmail(email string) (model.TicketDomainInterface, *rest_err.RestErr) {
	collection_name := os.Getenv(MONGODB_TICKET_COLLECTION)
	collection := tr.databaseConnection.Collection(collection_name)

	ticketEntity := &entity.TicketEntity{}

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

func (tr *ticketRepository) FindTicketById(id string) (model.TicketDomainInterface, *rest_err.RestErr) {
	collection_name := os.Getenv(MONGODB_TICKET_COLLECTION)
	collection := tr.databaseConnection.Collection(collection_name)

	ticketEntity := &entity.TicketEntity{}

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
