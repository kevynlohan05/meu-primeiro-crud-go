package service

import (
	"fmt"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) *rest_err.RestErr {

	userDomain.EncryptPassword()
	fmt.Println(userDomain.GetName())
	fmt.Println(userDomain.GetEmail())
	fmt.Println(userDomain.GetPassword())

	return nil
}
