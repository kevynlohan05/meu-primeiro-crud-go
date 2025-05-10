package service

import (
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	userModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user"
)

func (ud *userDomainService) FindUserByIdServices(id string) (userModel.UserDomainInterface, *rest_err.RestErr) {

	return ud.userRepository.FindUserById(id)
}

func (ud *userDomainService) FindUserByEmailServices(email string) (userModel.UserDomainInterface, *rest_err.RestErr) {

	return ud.userRepository.FindUserByEmail(email)
}

func (ud *userDomainService) findUserByEmailAndPasswordServices(email, password string) (userModel.UserDomainInterface, *rest_err.RestErr) {

	return ud.userRepository.FindUserByEmailAndPassword(email, password)
}
