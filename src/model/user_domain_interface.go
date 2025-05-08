package model

type UserDomainInterface interface {
	GetName() string
	GetEmail() string
	GetPassword() string
	GetDepartment() string
	GetRole() string

	GetID() string

	SetID(string)

	EncryptPassword()
}

func NewUserDomain(name, email, password, department, role string) UserDomainInterface {
	return &userDomain{
		name:       name,
		email:      email,
		password:   password,
		department: department,
		role:       role,
	}
}
