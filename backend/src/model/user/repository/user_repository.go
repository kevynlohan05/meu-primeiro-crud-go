package repository

import (
	"database/sql"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	userModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user"
)

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db,
	}
}

type userRepository struct {
	db *sql.DB
}

type UserRepository interface {
	CreateUser(userDomain userModel.UserDomainInterface) (userModel.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(userId string, userDomain userModel.UserDomainInterface) *rest_err.RestErr
	DeleteUser(userId string) *rest_err.RestErr
	FindUserByEmail(email string) (userModel.UserDomainInterface, *rest_err.RestErr)
	FindUserById(id string) (userModel.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmailAndPassword(email, password string) (userModel.UserDomainInterface, *rest_err.RestErr)
}
