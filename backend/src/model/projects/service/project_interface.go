package service

import (
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	projectModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/projects"
	repositoryProject "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/projects/repository"
)

func NewProjectService(projectRepository repositoryProject.ProjectRepository) ProjectDomainService {
	return &projectDomainService{projectRepository}
}

type projectDomainService struct {
	projectRepository repositoryProject.ProjectRepository
}

type ProjectDomainService interface {
	CreateProjectServices(projectModel.ProjectDomainInterface) (projectModel.ProjectDomainInterface, *rest_err.RestErr)
	UpdateProject(string, projectModel.ProjectDomainInterface) *rest_err.RestErr
	FindProjectByIdServices(id string) (projectModel.ProjectDomainInterface, *rest_err.RestErr)
	FindProjectByNameServices(name string) (projectModel.ProjectDomainInterface, *rest_err.RestErr)
	FindAllProjectsServices() ([]projectModel.ProjectDomainInterface, *rest_err.RestErr)
	FindProjectByAsanaIdServices(asanaId string) (projectModel.ProjectDomainInterface, *rest_err.RestErr)
	DeleteProject(string) *rest_err.RestErr
}
