package service

import (
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model/repository"
)

func NewUserDomainService(userRepository repository.UserRepository) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUserServices(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	FindUserByIdServices(email string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmailServices(email string) (model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
