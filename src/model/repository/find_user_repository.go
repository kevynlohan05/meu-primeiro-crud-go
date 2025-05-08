package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model/repository/entity"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (ur *userRepository) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}

	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User with email %s not found", email)
			return nil, rest_err.NewNotFoundError(errorMessage)
		}

		errorMessage := fmt.Sprintf("Error while trying to find user: %s", err.Error())
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	return converter.ConvertUserEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserById(id string) (model.UserDomainInterface, *rest_err.RestErr) {

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectId}}

	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User with id %s not found", id)
			return nil, rest_err.NewNotFoundError(errorMessage)
		}

		errorMessage := fmt.Sprintf("Error while trying to find user: %s", err.Error())
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	return converter.ConvertUserEntityToDomain(*userEntity), nil
}
