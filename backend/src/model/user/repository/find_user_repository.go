package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model/converter"
	userModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user"
	userEntity "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user/repository/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (ur *userRepository) FindUserByEmail(email string) (userModel.UserDomainInterface, *rest_err.RestErr) {

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &userEntity.UserEntity{}

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

func (ur *userRepository) FindUserById(id string) (userModel.UserDomainInterface, *rest_err.RestErr) {

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &userEntity.UserEntity{}

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

func (ur *userRepository) FindUserByEmailAndPassword(email, password string) (userModel.UserDomainInterface, *rest_err.RestErr) {

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &userEntity.UserEntity{}

	filter := bson.D{
		{Key: "email", Value: email},
		{Key: "password", Value: password},
	}

	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := ("User or passoword is invalid")
			return nil, rest_err.NewForbiddenError(errorMessage)
		}

		errorMessage := fmt.Sprintf("Error while trying to find user by email and password: %s", err.Error())
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	return converter.ConvertUserEntityToDomain(*userEntity), nil
}
