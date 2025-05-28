package repository

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model/converter"
	userModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user"
)

func (ur *userRepository) CreateUser(userDomain userModel.UserDomainInterface) (userModel.UserDomainInterface, *rest_err.RestErr) {
	log.Println("Converting domain to entity")
	value := converter.ConvertUserDomainToEntity(userDomain)

	if value == nil {
		log.Println("Error: Conversion to entity failed")
		return nil, rest_err.NewInternalServerError("Failed to convert domain to entity")
	}

	// Convertendo o campo Projects para JSON
	projectsJSON, err := json.Marshal(value.Projects)
	if err != nil {
		log.Println("Error marshaling projects:", err)
		return nil, rest_err.NewInternalServerError("Failed to encode projects to JSON")
	}

	query := `
	INSERT INTO users (name, email, password, phone, enterprise, department, projects, role)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?)
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
	)

	if err != nil {
		log.Println("Error inserting user into MySQL:", err)
		return nil, rest_err.NewInternalServerError(fmt.Sprintf("Erro ao inserir usuário: %v", err))
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		log.Println("Error retrieving inserted ID:", err)
		return nil, rest_err.NewInternalServerError("Erro ao obter ID do usuário")
	}

	log.Println("User inserted successfully into MySQL")

	value.ID = int(insertedID)

	log.Println("Converting entity back to domain")
	return converter.ConvertUserEntityToDomain(*value), nil
}
