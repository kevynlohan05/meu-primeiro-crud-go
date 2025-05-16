package repository

import (
	"context"
	"os"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	ticketModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (tr *ticketRepository) AddComment(ticketId string, comment ticketModel.CommentDomain) *rest_err.RestErr {
	collectionName := os.Getenv(MONGODB_TICKET_COLLECTION)
	collection := tr.databaseConnection.Collection(collectionName)

	objectID, err := primitive.ObjectIDFromHex(ticketId)
	if err != nil {
		return rest_err.NewBadRequestError("ID do ticket inválido")
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$push": bson.M{
			"comments": comment,
		},
	}

	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return rest_err.NewInternalServerError("Erro ao adicionar comentário")
	}

	return nil
}
