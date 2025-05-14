package repository

import (
	"context"
	"fmt"
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

func (tr *ticketRepository) FindAllTicketsByEmail(email string) ([]ticketModel.TicketDomainInterface, *rest_err.RestErr) {
	fmt.Printf("🔍 Buscando tickets para email exato: [%s]\n", email)
	collectionName := os.Getenv(MONGODB_TICKET_COLLECTION)
	collection := tr.databaseConnection.Collection(collectionName)

	filter := bson.M{"request_user": bson.M{"$eq": email}}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		fmt.Println("❌ Erro ao buscar tickets:", err.Error())
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	var tickets []ticketModel.TicketDomainInterface
	for cursor.Next(context.Background()) {
		var entity ticketEntity.TicketEntity
		if err := cursor.Decode(&entity); err != nil {
			fmt.Println("❌ Erro ao decodificar ticket:", err.Error())
			return nil, rest_err.NewInternalServerError(err.Error())
		}
		tickets = append(tickets, converter.ConvertTicketEntityToDomain(entity))
	}

	if len(tickets) == 0 {
		fmt.Println("⚠️ Nenhum ticket encontrado para:", email)
		return nil, rest_err.NewNotFoundError("Nenhum ticket encontrado para o e-mail informado")
	}

	return tickets, nil
}

func (tr *ticketRepository) FindAllTickets() ([]ticketModel.TicketDomainInterface, *rest_err.RestErr) {
	collectionName := os.Getenv(MONGODB_TICKET_COLLECTION)
	collection := tr.databaseConnection.Collection(collectionName)

	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		fmt.Println("❌ Erro ao buscar todos os tickets:", err.Error())
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	var tickets []ticketModel.TicketDomainInterface
	for cursor.Next(context.Background()) {
		var entity ticketEntity.TicketEntity
		if err := cursor.Decode(&entity); err != nil {
			fmt.Println("❌ Erro ao decodificar ticket:", err.Error())
			return nil, rest_err.NewInternalServerError(err.Error())
		}
		tickets = append(tickets, converter.ConvertTicketEntityToDomain(entity))
	}

	if len(tickets) == 0 {
		fmt.Println("⚠️ Nenhum ticket encontrado.")
		return nil, rest_err.NewNotFoundError("Nenhum ticket encontrado.")
	}

	return tickets, nil
}
