package service

import (
	"log"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	projectModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/projects"
)

func (pd *projectDomainService) CreateProjectServices(projectDomain projectModel.ProjectDomainInterface) (projectModel.ProjectDomainInterface, *rest_err.RestErr) {

	// Check if a project with the same name already exists
	existingProject, err := pd.FindProjectByNameServices(projectDomain.GetName())
	if err != nil {
		log.Printf("Error checking existing project by name: %v\n", err)
		return nil, err
	}
	if existingProject != nil {
		return nil, rest_err.NewBadRequestError("Project with this name already exists")
	}

	// Create the project in the repository
	createdProject, err := pd.projectRepository.CreateProject(projectDomain)
	if err != nil {
		log.Printf("Error creating project in repository: %v\n", err)
		return nil, err
	}

	if createdProject == nil {
		log.Println("Repository returned nil after creating project")
		return nil, rest_err.NewInternalServerError("Failed to create project in repository")
	}

	log.Printf("Project '%s' created successfully\n", projectDomain.GetName())
	return createdProject, nil
}
