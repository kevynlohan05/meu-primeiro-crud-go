package service

import (
	"log"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	projectModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/projects"
)

func (pd *projectDomainService) CreateProjectServices(projectDomain projectModel.ProjectDomainInterface) (projectModel.ProjectDomainInterface, *rest_err.RestErr) {

	project, _ := pd.FindProjectByNameServices(projectDomain.GetName())
	if project != nil {
		return nil, rest_err.NewBadRequestError("Project with this name already exists")
	}

	projectDomainRepository, err := pd.projectRepository.CreateProject(projectDomain)
	if err != nil {
		log.Println("Error in repository:", err)
		return nil, err
	}

	if projectDomainRepository == nil {
		log.Println("Error: projectDomainRepository is nil")
		return nil, rest_err.NewInternalServerError("Failed to create project in repository")
	}

	log.Println("Project created successfully")
	return projectDomainRepository, nil
}
