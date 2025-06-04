package repository

import (
	"errors"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model/converter"
	projectModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/projects"
)

// CreateProject inserts a new project into the database.
// It returns the created project domain object or an error if any occurs.
func (pr *projectRepository) CreateProject(projectDomain projectModel.ProjectDomainInterface) (projectModel.ProjectDomainInterface, *rest_err.RestErr) {
	log.Println("Starting conversion from domain to entity")
	entity := converter.ConvertProjectDomainToEntity(projectDomain)

	query := "INSERT INTO projects (name, asana_project_id) VALUES (?, ?)"

	// Execute the insert query
	result, err := pr.db.Exec(query,
		entity.Name,
		entity.IdAsana,
	)
	if err != nil {
		log.Printf("Error inserting project into DB: %v\n", err)

		// Check if the error is due to duplicate entry (MySQL error code 1062)
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return nil, rest_err.NewBadRequestValidationError("Error creating project", []rest_err.Causes{
				{Field: "name", Message: "A project with this name already exists"},
			})
		}

		return nil, rest_err.NewInternalServerError("Internal error when inserting project")
	}

	// Get the last inserted ID and assign it to the entity
	insertedID, err := result.LastInsertId()
	if err != nil {
		log.Printf("Failed to retrieve last insert ID: %v\n", err)
		return nil, rest_err.NewInternalServerError("Failed to get project ID after creation")
	}
	entity.ID = int(insertedID)

	log.Printf("Project created successfully with ID: %d\n", entity.ID)

	// Convert entity back to domain and return
	return converter.ConvertProjectEntityToDomain(*entity), nil
}
