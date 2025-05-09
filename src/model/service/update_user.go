package service

import (
	"log"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model"
)

func (ud *userDomainService) UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {

	log.Println("Calling repository to update user")

	err := ud.userRepository.UpdateUser(userId, userDomain)
	if err != nil {
		log.Println("Error in repository:", err)
		return err
	}

	log.Println("User update successfully")

	return nil
}
