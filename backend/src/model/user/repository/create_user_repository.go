package repository

import (
	"database/sql"
	"errors"
	"log"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model/converter"
	userModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user"
)

func (ur *userRepository) CreateUser(userDomain userModel.UserDomainInterface) (userModel.UserDomainInterface, *rest_err.RestErr) {
	log.Println("Converting domain to entity")
	value := converter.ConvertUserDomainToEntity(userDomain)

	projectIDs := make([]int, 0, len(userDomain.GetProjects()))
	for _, projectName := range userDomain.GetProjects() {
		var projectID int
		err := ur.db.QueryRow("SELECT id FROM projects WHERE name = ?", projectName).Scan(&projectID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				log.Printf("Project '%s' not found, aborting user creation\n", projectName)
				return nil, rest_err.NewBadRequestError("Project '" + projectName + "' does not exist")
			}
			log.Println("Error checking project existence:", err)
			return nil, rest_err.NewInternalServerError("Error while validating project list")
		}
		projectIDs = append(projectIDs, projectID)
	}

	query := `
		INSERT INTO users (name, email, password, phone, enterprise, department, role)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	result, err := ur.db.Exec(query,
		value.Name,
		value.Email,
		value.Password,
		value.Phone,
		value.Enterprise,
		value.Department,
		value.Role,
	)
	if err != nil {
		log.Println("Error inserting user:", err)
		return nil, rest_err.NewInternalServerError("Error while inserting user")
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		log.Println("Error retrieving last inserted user ID:", err)
		return nil, rest_err.NewInternalServerError("Error while retrieving user ID")
	}
	value.ID = int(insertedID)

	for _, projectID := range projectIDs {
		_, err = ur.db.Exec("INSERT INTO user_projects (user_id, project_id) VALUES (?, ?)", insertedID, projectID)
		if err != nil {
			log.Printf("Error associating user with project ID %d: %v\n", projectID, err)
			continue
		}
	}

	log.Println("User and associated projects created successfully")
	return converter.ConvertUserEntityToDomain(*value, userDomain.GetProjects()), nil
}
