package repository

import (
	"errors"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model/converter"
	projectModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/projects"
)

func (pr *projectRepository) CreateProject(projectDomain projectModel.ProjectDomainInterface) (projectModel.ProjectDomainInterface, *rest_err.RestErr) {
	log.Println("Converting domain to entity")
	value := converter.ConvertProjectDomainToEntity(projectDomain)

	query := "INSERT INTO projects (name, asana_project_id) VALUES (?, ?)"

	result, err := pr.db.Exec(query,
		value.Name,
		value.IdAsana,
	)
	if err != nil {
		log.Println("Error inserting project:", err)

		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return nil, rest_err.NewBadRequestValidationError("Erro ao criar projeto", []rest_err.Causes{
				{Field: "name", Message: "Já existe um projeto com esse nome"},
			})
		}

		return nil, rest_err.NewInternalServerError("Error inserting project")
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		return nil, rest_err.NewInternalServerError("Erro ao obter ID do project")
	}
	value.ID = int(insertedID)

	log.Println("Project created successfully with ID:", value.ID)
	return converter.ConvertProjectEntityToDomain(*value), nil
}
