package service

import (
	"log"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model"
)

func (ud *userDomainService) LoginUserServices(userDomain model.UserDomainInterface) (model.UserDomainInterface, string, *rest_err.RestErr) {

	log.Println("Encrypting password")
	userDomain.EncryptPassword()

	user, err := ud.findUserByEmailAndPasswordServices(userDomain.GetEmail(), userDomain.GetPassword())
	if err != nil {
		log.Println("User already exists")
		return nil, "", err
	}

	token, err := user.GenerateToken()
	if err != nil {
		return nil, "", err
	}

	log.Println("Login user successfully")

	return user, token, nil
}
