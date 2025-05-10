package service

import (
	"log"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	userModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user"
)

func (ud *userDomainService) LoginUserServices(userDomain userModel.UserDomainInterface) (userModel.UserDomainInterface, string, *rest_err.RestErr) {

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
