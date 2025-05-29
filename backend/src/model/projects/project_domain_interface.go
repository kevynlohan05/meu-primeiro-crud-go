package projects

type ProjectDomainInterface interface {
	GetID() string
	GetName() string
	GetIdAsana() string
	SetID(string)
}

func NewProjectDomain(name, idAsana string) ProjectDomainInterface {
	return &Project{
		name:    name,
		idAsana: idAsana,
	}
}

func NewProjectUpdateDomain(name, idAsana string) ProjectDomainInterface {
	return &Project{
		name:    name,
		idAsana: idAsana,
	}
}
