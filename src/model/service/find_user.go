package service

import (
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model"
)

func (ud *userDomainService) FindUserByIdServices(id string) (model.UserDomainInterface, *rest_err.RestErr) {

	return ud.userRepository.FindUserById(id)
}

func (ud *userDomainService) FindUserByEmailServices(email string) (model.UserDomainInterface, *rest_err.RestErr) {

	return ud.userRepository.FindUserByEmail(email)
}
