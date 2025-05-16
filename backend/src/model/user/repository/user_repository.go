package repository

import (
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	userModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"
)

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{
		database,
	}

}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(userDomain userModel.UserDomainInterface) (userModel.UserDomainInterface, *rest_err.RestErr)

	UpdateUser(userId string, userDomain userModel.UserDomainInterface) *rest_err.RestErr

	DeleteUser(userId string) *rest_err.RestErr

	FindUserByEmail(email string) (userModel.UserDomainInterface, *rest_err.RestErr)

	FindUserById(id string) (userModel.UserDomainInterface, *rest_err.RestErr)

	FindUserByEmailAndPassword(email, password string) (userModel.UserDomainInterface, *rest_err.RestErr)
}
