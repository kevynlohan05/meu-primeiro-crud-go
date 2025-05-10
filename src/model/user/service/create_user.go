package service

import (
	"log"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	userModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user"
)

func (ud *userDomainService) CreateUserServices(userDomain userModel.UserDomainInterface) (userModel.UserDomainInterface, *rest_err.RestErr) {

	user, _ := ud.FindUserByEmailServices(userDomain.GetEmail())
	if user != nil {
		log.Println("User already exists")
		return nil, rest_err.NewBadRequestError("Email already registered")
	}

	log.Println("Encrypting password")
	userDomain.EncryptPassword()

	log.Println("Calling repository to create user")
	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		log.Println("Error in repository:", err)
		return nil, err
	}

	if userDomainRepository == nil {
		log.Println("Error: userDomainRepository is nil")
		return nil, rest_err.NewInternalServerError("Failed to create user in repository")
	}

	log.Println("User created successfully")

	return userDomainRepository, nil
}
