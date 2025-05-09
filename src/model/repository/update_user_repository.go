package repository

import (
	"context"
	"log"
	"os"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ur *userRepository) UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	log.Println("Converting domain to entity")

	value := converter.ConvertUserDomainToEntity(userDomain)
	userIdHex, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.D{{Key: "_id", Value: userIdHex}}
	update := bson.D{{Key: "$set", Value: value}}

	log.Println("Update user into MongoDB")
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println("Error Update user into MongoDB:", err)
		return rest_err.NewInternalServerError(err.Error())

	}

	log.Println("Update user successfully into MongoDB")
	return nil
}
