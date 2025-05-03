package model

import (
	"fmt"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
)

func (ud *userDomain) CreateUser() *rest_err.RestErr {

	ud.EncryptPassword()
	fmt.Println(ud)

	return nil
}
