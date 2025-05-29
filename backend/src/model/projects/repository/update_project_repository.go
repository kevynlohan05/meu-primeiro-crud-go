package repository

import (
	"log"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model/converter"
	projectModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/projects"
)

func (pr *projectRepository) UpdateProject(projectId string, projectDomain projectModel.ProjectDomainInterface) *rest_err.RestErr {
	log.Println("Converting domain to entity")
	value := converter.ConvertProjectDomainToEntity(projectDomain)
	if value == nil {
		log.Println("Error: Conversion to entity failed")
		return rest_err.NewInternalServerError("Failed to convert domain to entity")
	}
	query := `
		UPDATE projects SET
			name = ?,
			asana_project_id = ?
		WHERE id = ?
	`
	result, err := pr.db.Exec(query,
		value.Name,
		value.IdAsana,
		projectId,
	)

	if err != nil {
		log.Println("Error updating project in MySQL:", err)
		return rest_err.NewInternalServerError("Error updating project: " + err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error getting affected rows:", err)
		return rest_err.NewInternalServerError("Error verifying update operation: " + err.Error())
	}
	if rowsAffected == 0 {
		log.Println("No project found with id:", projectId)
		return rest_err.NewNotFoundError("Project with id " + projectId + " not found")
	}
	log.Println("Project updated successfully")
	return nil
}
