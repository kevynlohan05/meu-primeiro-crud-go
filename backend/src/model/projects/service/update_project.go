package service

import (
	"log"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	projectModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/projects"
)

func (pd *projectDomainService) UpdateProject(projectId string, projectDomain projectModel.ProjectDomainInterface) *rest_err.RestErr {

	log.Printf("Attempting to update project with ID: %s\n", projectId)

	err := pd.projectRepository.UpdateProject(projectId, projectDomain)
	if err != nil {
		log.Printf("Repository error while updating project with ID %s: %v\n", projectId, err)
		return err
	}

	log.Printf("Project with ID %s updated successfully\n", projectId)

	return nil
}
