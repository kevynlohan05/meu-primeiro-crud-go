package repository

import (
	"log"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model/converter"
	userModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user"
)

func (ur *userRepository) CreateUser(userDomain userModel.UserDomainInterface) (userModel.UserDomainInterface, *rest_err.RestErr) {
	log.Println("Converting domain to entity")
	value := converter.ConvertUserDomainToEntity(userDomain)

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

	for _, projectName := range userDomain.GetProjects() {
		var projectID int

		// Check if project already exists
		err := ur.db.QueryRow("SELECT id FROM projects WHERE name = ?", projectName).Scan(&projectID)
		if err != nil {
			// Create project if not exists
			res, err := ur.db.Exec("INSERT INTO projects (name) VALUES (?)", projectName)
			if err != nil {
				log.Println("Error inserting project:", err)
				continue
			}
			lastID, err := res.LastInsertId()
			if err != nil {
				log.Println("Error retrieving project ID:", err)
				continue
			}
			projectID = int(lastID)
		}

		// Associate user with project
		_, err = ur.db.Exec("INSERT INTO user_projects (user_id, project_id) VALUES (?, ?)", insertedID, projectID)
		if err != nil {
			log.Println("Error inserting into user_projects:", err)
			continue
		}
	}

	log.Println("User and associated projects created successfully")
	return converter.ConvertUserEntityToDomain(*value, userDomain.GetProjects()), nil
}
