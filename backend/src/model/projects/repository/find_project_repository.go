package repository

import (
	"database/sql"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model/converter"
	projectModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/projects"
	ProjectEntity "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/projects/repository/entity"
)

func (pr *projectRepository) FindProjectById(id string) (projectModel.ProjectDomainInterface, *rest_err.RestErr) {
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
			return nil, rest_err.NewNotFoundError("Project with ID " + id + " not found")
		}
		return nil, rest_err.NewInternalServerError("Error finding project: " + err.Error())
	}
	return converter.ConvertProjectEntityToDomain(entity), nil
}

func (pr *projectRepository) FindProjectByName(name string) (projectModel.ProjectDomainInterface, *rest_err.RestErr) {
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
			return nil, rest_err.NewNotFoundError("Project with name " + name + " not found")
		}
		return nil, rest_err.NewInternalServerError("Error finding project: " + err.Error())
	}
	return converter.ConvertProjectEntityToDomain(entity), nil
}

func (pr *projectRepository) FindAllProjects() ([]projectModel.ProjectDomainInterface, *rest_err.RestErr) {
	query := `
		SELECT id, name, asana_project_id
		FROM projects
	`
	rows, err := pr.db.Query(query)
	if err != nil {
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
			return nil, rest_err.NewInternalServerError("Error scanning project: " + err.Error())
		}
		project := converter.ConvertProjectEntityToDomain(entity)
		if project != nil {
			projects = append(projects, project)
		}
	}
	if err := rows.Err(); err != nil {
		return nil, rest_err.NewInternalServerError("Error iterating over projects: " + err.Error())
	}
	if len(projects) == 0 {
		return nil, rest_err.NewNotFoundError("No projects found")
	}
	return projects, nil
}

func (pr *projectRepository) FindProjectByAsanaId(asanaId string) (projectModel.ProjectDomainInterface, *rest_err.RestErr) {
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
			return nil, rest_err.NewNotFoundError("Project with Asana ID " + asanaId + " not found")
		}
		return nil, rest_err.NewInternalServerError("Error finding project: " + err.Error())
	}
	return converter.ConvertProjectEntityToDomain(entity), nil
}
