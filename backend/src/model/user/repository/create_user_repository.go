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
		return nil, rest_err.NewInternalServerError("Erro ao inserir usuário")
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		return nil, rest_err.NewInternalServerError("Erro ao obter ID do usuário")
	}
	value.ID = int(insertedID)

	for _, projectName := range userDomain.GetProjects() {
		var projectID int

		// Verifica se projeto existe
		err := ur.db.QueryRow("SELECT id FROM projects WHERE name = ?", projectName).Scan(&projectID)
		if err != nil {
			// Cria se não existe
			res, err := ur.db.Exec("INSERT INTO projects (name) VALUES (?)", projectName)
			if err != nil {
				log.Println("Erro ao inserir projeto:", err)
				continue
			}
			lastID, _ := res.LastInsertId()
			projectID = int(lastID)
		}

		// Relaciona com usuário
		_, err = ur.db.Exec("INSERT INTO user_projects (user_id, project_id) VALUES (?, ?)", insertedID, projectID)
		if err != nil {
			log.Println("Erro ao inserir user_project:", err)
			continue
		}
	}

	log.Println("Usuário e projetos inseridos com sucesso")
	return converter.ConvertUserEntityToDomain(*value, userDomain.GetProjects()), nil
}
