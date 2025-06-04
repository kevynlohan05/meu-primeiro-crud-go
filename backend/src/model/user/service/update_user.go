package service

import (
	"log"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	userModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user"
)

// UpdateUser updates an existing user with the given ID and domain data.
func (ud *userDomainService) UpdateUser(userId string, userDomain userModel.UserDomainInterface) *rest_err.RestErr {
	log.Println("Calling repository to update user")

	err := ud.userRepository.UpdateUser(userId, userDomain)
	if err != nil {
		log.Println("Error updating user in repository:", err)
		return err
	}

	log.Println("User updated successfully")
	return nil
}
