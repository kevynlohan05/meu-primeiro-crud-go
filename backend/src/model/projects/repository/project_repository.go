package repository

import (
	"database/sql"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	projectModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/projects"
)

func NewProjectRepository(db *sql.DB) ProjectRepository {
	return &projectRepository{
		db,
	}
}

type projectRepository struct {
	db *sql.DB
}

type ProjectRepository interface {
	CreateProject(projectDomain projectModel.ProjectDomainInterface) (projectModel.ProjectDomainInterface, *rest_err.RestErr)
	UpdateProject(projectId string, projectDomain projectModel.ProjectDomainInterface) *rest_err.RestErr
	DeleteProject(projectId string) *rest_err.RestErr
	FindProjectById(id string) (projectModel.ProjectDomainInterface, *rest_err.RestErr)
	FindProjectByName(name string) (projectModel.ProjectDomainInterface, *rest_err.RestErr)
}
