package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
)

func NewUserDomain(name, email, password, department, role string) UserDomainInterface {
	return &userDomain{name, email, password, department, role}
}

type userDomain struct {
	Name       string
	Email      string
	Password   string
	Department string
	Role       string
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *rest_err.RestErr
	UpdateUser(string) *rest_err.RestErr
	FindUser(string) (*userDomain, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
