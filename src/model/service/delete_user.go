package service

import (
	"log"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
)

func (ud *userDomainService) DeleteUser(userId string) *rest_err.RestErr {
	log.Println("Calling repository to delete user")

	err := ud.userRepository.DeleteUser(userId)
	if err != nil {
		log.Println("Error in repository:", err)
		return err
	}

	log.Println("User delete successfully")

	return nil
}
