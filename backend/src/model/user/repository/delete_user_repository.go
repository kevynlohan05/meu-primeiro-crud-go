package repository

import (
	"context"
	"log"
	"os"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ur *userRepository) DeleteUser(userId string) *rest_err.RestErr {
	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	log.Println("Converting domain to entity")

	userIdHex, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.D{{Key: "_id", Value: userIdHex}}

	log.Println("Delete user into MongoDB")
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Println("Error Delete user into MongoDB:", err)
		return rest_err.NewInternalServerError(err.Error())

	}

	log.Println("Delete user successfully into MongoDB")
	return nil
}
