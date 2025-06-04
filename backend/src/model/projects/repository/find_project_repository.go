package repository

import (
	"database/sql"
	"log"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model/converter"
	projectModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/projects"
	ProjectEntity "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/projects/repository/entity"
)

// FindProjectById retrieves a project from the database by its ID.
func (pr *projectRepository) FindProjectById(id string) (projectModel.ProjectDomainInterface, *rest_err.RestErr) {
	log.Printf("Searching for project by ID: %s\n", id)

	query := `
		SELECT id, name, asana_project_id
		FROM projects WHERE id = ?
	`
	row := pr.db.QueryRow(query, id)
	var entity ProjectEntity.ProjectEntity
	err := row.Scan(
		&entity.ID,
		&entity.Name,
		&entity.IdAsana,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Project not found with ID: %s\n", id)
			return nil, rest_err.NewNotFoundError("Project with ID " + id + " not found")
		}
		log.Printf("Error retrieving project by ID %s: %v\n", id, err)
		return nil, rest_err.NewInternalServerError("Error finding project: " + err.Error())
	}

	log.Printf("Project found with ID: %s\n", id)
	return converter.ConvertProjectEntityToDomain(entity), nil
}

// FindProjectByName retrieves a project from the database by its name.
func (pr *projectRepository) FindProjectByName(name string) (projectModel.ProjectDomainInterface, *rest_err.RestErr) {
	log.Printf("Searching for project by name: %s\n", name)

	query := `
		SELECT id, name, asana_project_id
		FROM projects WHERE name = ?
	`
	row := pr.db.QueryRow(query, name)
	var entity ProjectEntity.ProjectEntity
	err := row.Scan(
		&entity.ID,
		&entity.Name,
		&entity.IdAsana,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Project not found with name: %s\n", name)
			return nil, rest_err.NewNotFoundError("Project with name " + name + " not found")
		}
		log.Printf("Error retrieving project by name %s: %v\n", name, err)
		return nil, rest_err.NewInternalServerError("Error finding project: " + err.Error())
	}

	log.Printf("Project found with name: %s\n", name)
	return converter.ConvertProjectEntityToDomain(entity), nil
}

// FindAllProjects retrieves all projects from the database.
func (pr *projectRepository) FindAllProjects() ([]projectModel.ProjectDomainInterface, *rest_err.RestErr) {
	log.Println("Retrieving all projects")

	query := `
		SELECT id, name, asana_project_id
		FROM projects
	`
	rows, err := pr.db.Query(query)
	if err != nil {
		log.Printf("Error querying projects: %v\n", err)
		return nil, rest_err.NewInternalServerError("Error finding projects: " + err.Error())
	}
	defer rows.Close()

	var projects []projectModel.ProjectDomainInterface
	for rows.Next() {
		var entity ProjectEntity.ProjectEntity
		if err := rows.Scan(
			&entity.ID,
			&entity.Name,
			&entity.IdAsana,
		); err != nil {
			log.Printf("Error scanning project row: %v\n", err)
			return nil, rest_err.NewInternalServerError("Error scanning project: " + err.Error())
		}
		project := converter.ConvertProjectEntityToDomain(entity)
		if project != nil {
			projects = append(projects, project)
		}
	}
	if err := rows.Err(); err != nil {
		log.Printf("Row iteration error: %v\n", err)
		return nil, rest_err.NewInternalServerError("Error iterating over projects: " + err.Error())
	}

	if len(projects) == 0 {
		log.Println("No projects found in database")
		return nil, rest_err.NewNotFoundError("No projects found")
	}

	log.Printf("Total projects found: %d\n", len(projects))
	return projects, nil
}

// FindProjectByAsanaId retrieves a project from the database by its Asana project ID.
func (pr *projectRepository) FindProjectByAsanaId(asanaId string) (projectModel.ProjectDomainInterface, *rest_err.RestErr) {
	log.Printf("Searching for project by Asana ID: %s\n", asanaId)

	query := `
		SELECT id, name, asana_project_id
		FROM projects WHERE asana_project_id = ?
	`
	row := pr.db.QueryRow(query, asanaId)
	var entity ProjectEntity.ProjectEntity
	err := row.Scan(
		&entity.ID,
		&entity.Name,
		&entity.IdAsana,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Project not found with Asana ID: %s\n", asanaId)
			return nil, rest_err.NewNotFoundError("Project with Asana ID " + asanaId + " not found")
		}
		log.Printf("Error retrieving project by Asana ID %s: %v\n", asanaId, err)
		return nil, rest_err.NewInternalServerError("Error finding project: " + err.Error())
	}

	log.Printf("Project found with Asana ID: %s\n", asanaId)
	return converter.ConvertProjectEntityToDomain(entity), nil
}
