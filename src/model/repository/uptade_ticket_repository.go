package repository

import (
	"context"
	"log"
	"os"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (tr *ticketRepository) UpdateTicket(ticketId string, ticketDomain model.TicketDomainInterface) *rest_err.RestErr {
	collection_name := os.Getenv(MONGODB_TICKET_COLLECTION)
	collection := tr.databaseConnection.Collection(collection_name)

	log.Println("Converting domain to entity")

	value := converter.ConvertTicketDomainToEntity(ticketDomain)
	ticketIdHex, _ := primitive.ObjectIDFromHex(ticketId)

	filter := primitive.D{{Key: "_id", Value: ticketIdHex}}
	update := primitive.D{{Key: "$set", Value: value}}

	log.Println("Update ticket into MongoDB")
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println("Error Update ticket into MongoDB:", err)
		return rest_err.NewInternalServerError(err.Error())

	}

	log.Println("Update ticket successfully into MongoDB")
	return nil
}
