package repository

import (
	"context"
	"log"
	"os"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (tr *ticketRepository) DeleteTicket(ticketId string) *rest_err.RestErr {
	collection_name := os.Getenv(MONGODB_TICKET_COLLECTION)
	collection := tr.databaseConnection.Collection(collection_name)

	log.Println("Converting domain to entity")

	ticketIdHex, _ := primitive.ObjectIDFromHex(ticketId)

	filter := primitive.D{{Key: "_id", Value: ticketIdHex}}

	log.Println("Delete ticket into MongoDB")
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Println("Error Delete ticket into MongoDB:", err)
		return rest_err.NewInternalServerError(err.Error())

	}

	log.Println("Delete ticket successfully into MongoDB")
	return nil
}
