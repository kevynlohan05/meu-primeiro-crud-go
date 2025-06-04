package service

import (
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	userModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user"
)

// FindUserByIdServices retrieves a user by their ID.
func (ud *userDomainService) FindUserByIdServices(id string) (userModel.UserDomainInterface, *rest_err.RestErr) {
	return ud.userRepository.FindUserById(id)
}

// FindUserByEmailServices retrieves a user by their email.
func (ud *userDomainService) FindUserByEmailServices(email string) (userModel.UserDomainInterface, *rest_err.RestErr) {
	return ud.userRepository.FindUserByEmail(email)
}

// FindUserByEmailAndPasswordServices retrieves a user by email and password.
func (ud *userDomainService) findUserByEmailAndPasswordServices(email, password string) (userModel.UserDomainInterface, *rest_err.RestErr) {
	return ud.userRepository.FindUserByEmailAndPassword(email, password)
}
