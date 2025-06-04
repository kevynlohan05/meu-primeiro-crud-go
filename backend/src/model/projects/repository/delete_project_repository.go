package repository

import (
	"log"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
)

// DeleteProject removes a project from the database by its ID.
// Returns an error if the project does not exist or if any issue occurs during deletion.
func (pr *projectRepository) DeleteProject(projectId string) *rest_err.RestErr {
	log.Printf("Attempting to delete project with ID: %s\n", projectId)

	query := `DELETE FROM projects WHERE id = ?`

	result, err := pr.db.Exec(query, projectId)
	if err != nil {
		log.Printf("Error executing delete query: %v\n", err)
		return rest_err.NewInternalServerError("Error deleting project: " + err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error retrieving affected rows count: %v\n", err)
		return rest_err.NewInternalServerError("Error verifying delete operation: " + err.Error())
	}

	if rowsAffected == 0 {
		log.Printf("No project found with ID: %s\n", projectId)
		return rest_err.NewNotFoundError("Project with ID " + projectId + " not found")
	}

	log.Printf("Project with ID %s deleted successfully\n", projectId)
	return nil
}
