package repository

import (
	"fmt"
	"log"
	"encoding/json"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model/converter"
	userModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user"
)

func (ur *userRepository) UpdateUser(userId string, userDomain userModel.UserDomainInterface) *rest_err.RestErr {
	log.Println("Converting domain to entity")
	value := converter.ConvertUserDomainToEntity(userDomain)
	if value == nil {
		log.Println("Error: Conversion to entity failed")
		return rest_err.NewInternalServerError("Failed to convert domain to entity")
	}

	projectsJSON, err := json.Marshal(value.Projects)
	if err != nil {
		log.Println("Error marshalling projects:", err)
		return rest_err.NewInternalServerError("Failed to serialize user projects")
	}

	query := `
		UPDATE users SET
			name = ?,
			email = ?,
			password = ?,
			phone = ?,
			enterprise = ?,
			department = ?,
			projects = ?,
			role = ?
		WHERE id = ?
	`

	result, err := ur.db.Exec(query,
		value.Name,
		value.Email,
		value.Password,
		value.Phone,
		value.Enterprise,
		value.Department,
		string(projectsJSON),
		value.Role,
		userId,
	)

	if err != nil {
		log.Println("Error updating user in MySQL:", err)
		return rest_err.NewInternalServerError(fmt.Sprintf("Error updating user: %v", err))
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error getting affected rows:", err)
		return rest_err.NewInternalServerError("Error verifying update operation")
	}

	if rowsAffected == 0 {
		log.Println("No user found with id:", userId)
		return rest_err.NewNotFoundError(fmt.Sprintf("User with id %s not found", userId))
	}

	log.Println("User updated successfully in MySQL")
	return nil
}
