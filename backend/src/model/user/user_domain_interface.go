package model

import "github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"

type UserDomainInterface interface {
	GetName() string
	GetEmail() string
	GetPassword() string
	GetDepartment() string
	GetProjects() []string
	GetRole() string

	GetID() string

	SetID(string)
	SetProjects([]string)

	AddProject(string)

	EncryptPassword()
	GenerateToken() (string, *rest_err.RestErr)
}

func NewUserLoginDomain(email, password string) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
	}
}

func NewUserDomain(name, email, password, department, role string, projects []string) UserDomainInterface {
	return &userDomain{
		name:       name,
		email:      email,
		password:   password,
		department: department,
		role:       role,
		projects:   projects,
	}
}

func NewUserUpdateDomain(name, department string) UserDomainInterface {
	return &userDomain{
		name:       name,
		department: department,
	}
}
