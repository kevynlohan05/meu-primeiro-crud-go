package service

import (
	"log"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	userModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user"
)

func (ud *userDomainService) UpdateUser(userId string, userDomain userModel.UserDomainInterface) *rest_err.RestErr {

	log.Println("Calling repository to update user")

	err := ud.userRepository.UpdateUser(userId, userDomain)
	if err != nil {
		log.Println("Error in repository:", err)
		return err
	}

	log.Println("User update successfully")

	return nil
}
