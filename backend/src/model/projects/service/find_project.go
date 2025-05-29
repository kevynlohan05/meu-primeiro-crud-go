package service

import (
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	projectModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/projects"
)

func (pd *projectDomainService) FindProjectByIdServices(id string) (projectModel.ProjectDomainInterface, *rest_err.RestErr) {
	project, err := pd.projectRepository.FindProjectById(id)
	if err != nil {
		return nil, err
	}

	if project == nil {
		return nil, rest_err.NewNotFoundError("Project not found") // Return a not found error if project is nil
	}

	return project, nil
}

func (pd *projectDomainService) FindProjectByNameServices(name string) (projectModel.ProjectDomainInterface, *rest_err.RestErr) {

	return pd.projectRepository.FindProjectByName(name)
}
