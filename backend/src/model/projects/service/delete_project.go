package service

import (
	"log"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
)

func (pd *projectDomainService) DeleteProject(projectId string) *rest_err.RestErr {
	log.Println("Calling repository to delete project")

	err := pd.projectRepository.DeleteProject(projectId)
	if err != nil {
		log.Printf("Error deleting project in repository: %v\n", err)
		return err
	}

	log.Println("Project deleted successfully")
	return nil
}
