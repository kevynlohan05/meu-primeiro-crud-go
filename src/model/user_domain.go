package model

type userDomain struct {
	iD         string
	name       string
	email      string
	password   string
	department string
	role       string
}

func (ud *userDomain) GetID() string {
	return ud.iD
}

func (ud *userDomain) GetName() string {
	return ud.name
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}
func (ud *userDomain) GetPassword() string {
	return ud.password
}

func (ud *userDomain) GetDepartment() string {
	return ud.department
}

func (ud *userDomain) GetRole() string {
	return ud.role
}

func (ud *userDomain) SetID(id string) {
	ud.iD = id
}
