package repository

import (
	"log"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
)

func (pr *projectRepository) DeleteProject(projectId string) *rest_err.RestErr {
	log.Println("Deleting project in MySQL with ID:", projectId)

	query := `DELETE FROM projects WHERE id = ?`

	result, err := pr.db.Exec(query, projectId)
	if err != nil {
		log.Println("Error deleting project from MySQL:", err)
		return rest_err.NewInternalServerError("Error deleting project: " + err.Error())
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error getting affected rows:", err)
		return rest_err.NewInternalServerError("Error verifying delete operation: " + err.Error())
	}
	if rowsAffected == 0 {
		log.Println("No project found with ID:", projectId)
		return rest_err.NewNotFoundError("Project with ID " + projectId + " not found")
	}
	log.Println("Project deleted successfully from MySQL")
	return nil

}
