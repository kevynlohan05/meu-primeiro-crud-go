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
		// Return a not found error if no project with given ID exists
		return nil, rest_err.NewNotFoundError("Project not found")
	}

	return project, nil
}

func (pd *projectDomainService) FindProjectByNameServices(name string) (projectModel.ProjectDomainInterface, *rest_err.RestErr) {
	return pd.projectRepository.FindProjectByName(name)
}

func (pd *projectDomainService) FindAllProjectsServices() ([]projectModel.ProjectDomainInterface, *rest_err.RestErr) {
	projects, err := pd.projectRepository.FindAllProjects()
	if err != nil {
		return nil, err
	}

	if len(projects) == 0 {
		// Return not found error if there are no projects in the repository
		return nil, rest_err.NewNotFoundError("No projects found")
	}

	return projects, nil
}

func (pd *projectDomainService) FindProjectByAsanaIdServices(asanaId string) (projectModel.ProjectDomainInterface, *rest_err.RestErr) {
	project, err := pd.projectRepository.FindProjectByAsanaId(asanaId)
	if err != nil {
		return nil, err
	}

	if project == nil {
		return nil, rest_err.NewNotFoundError("Project with Asana ID " + asanaId + " not found")
	}

	return project, nil
}
