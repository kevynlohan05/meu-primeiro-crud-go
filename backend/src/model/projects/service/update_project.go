package service

import (
	"log"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	projectModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/projects"
)

func (pd *projectDomainService) UpdateProject(projectId string, projectDomain projectModel.ProjectDomainInterface) *rest_err.RestErr {

	log.Println("Calling repository to update project")

	err := pd.projectRepository.UpdateProject(projectId, projectDomain)
	if err != nil {
		log.Println("Error in repository:", err)
		return err
	}

	log.Println("Project updated successfully")

	return nil
}
