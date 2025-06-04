package service

import (
	"log"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	userModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user"
)

// LoginUserServices authenticates the user and generates a JWT token.
func (ud *userDomainService) LoginUserServices(userDomain userModel.UserDomainInterface) (userModel.UserDomainInterface, string, *rest_err.RestErr) {
	log.Println("Encrypting password")
	userDomain.EncryptPassword()

	log.Println("Verifying user credentials")
	user, err := ud.findUserByEmailAndPasswordServices(userDomain.GetEmail(), userDomain.GetPassword())
	if err != nil {
		log.Println("Invalid credentials")
		return nil, "", err
	}

	log.Println("Generating JWT token")
	token, err := user.GenerateToken()
	if err != nil {
		log.Println("Error generating token:", err)
		return nil, "", err
	}

	log.Println("User login successful")
	return user, token, nil
}
