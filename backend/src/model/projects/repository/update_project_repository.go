package repository

import (
	"log"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model/converter"
	projectModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/projects"
)

// UpdateProject updates the project identified by projectId with the provided projectDomain data.
func (pr *projectRepository) UpdateProject(projectId string, projectDomain projectModel.ProjectDomainInterface) *rest_err.RestErr {
	log.Println("Converting project domain to entity for update")
	value := converter.ConvertProjectDomainToEntity(projectDomain)
	if value == nil {
		log.Println("Failed to convert project domain to entity")
		return rest_err.NewInternalServerError("Failed to convert domain to entity")
	}

	query := `
		UPDATE projects SET
			name = ?,
			asana_project_id = ?
		WHERE id = ?
	`

	log.Printf("Executing update query for project ID: %s\n", projectId)
	result, err := pr.db.Exec(query,
		value.Name,
		value.IdAsana,
		projectId,
	)
	if err != nil {
		log.Printf("Error updating project in database: %v\n", err)
		return rest_err.NewInternalServerError("Error updating project: " + err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error fetching rows affected after update: %v\n", err)
		return rest_err.NewInternalServerError("Error verifying update operation: " + err.Error())
	}

	if rowsAffected == 0 {
		log.Printf("No project found to update with ID: %s\n", projectId)
		return rest_err.NewNotFoundError("Project with id " + projectId + " not found")
	}

	log.Printf("Project with ID %s updated successfully\n", projectId)
	return nil
}
