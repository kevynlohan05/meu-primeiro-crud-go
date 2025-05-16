package service

import (
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	userModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user"
	repositoryUser "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user/repository"
)

func NewUserDomainService(userRepository repositoryUser.UserRepository) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repositoryUser.UserRepository
}

type UserDomainService interface {
	CreateUserServices(userModel.UserDomainInterface) (userModel.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(string, userModel.UserDomainInterface) *rest_err.RestErr
	FindUserByIdServices(email string) (userModel.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmailServices(email string) (userModel.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
	LoginUserServices(userDomain userModel.UserDomainInterface) (userModel.UserDomainInterface, string, *rest_err.RestErr)
}
