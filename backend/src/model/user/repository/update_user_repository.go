package repository

import (
	"fmt"
	"log"

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

	query := `
		UPDATE users SET
			name = ?,
			email = ?,
			password = ?,
			phone = ?,
			enterprise = ?,
			department = ?,
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

	log.Println("Updating user_projects")

	_, err = ur.db.Exec("DELETE FROM user_projects WHERE user_id = ?", userId)
	if err != nil {
		log.Println("Erro ao remover projetos antigos:", err)
		return rest_err.NewInternalServerError("Erro ao atualizar os projetos do usuário")
	}

	for _, projectName := range userDomain.GetProjects() {
		var projectID int

		err := ur.db.QueryRow("SELECT id FROM projects WHERE name = ?", projectName).Scan(&projectID)
		if err != nil {
			res, err := ur.db.Exec("INSERT INTO projects (name) VALUES (?)", projectName)
			if err != nil {
				log.Println("Erro ao inserir novo projeto:", err)
				continue
			}
			lastID, _ := res.LastInsertId()
			projectID = int(lastID)
		}

		_, err = ur.db.Exec("INSERT INTO user_projects (user_id, project_id) VALUES (?, ?)", userId, projectID)
		if err != nil {
			log.Println("Erro ao relacionar projeto com usuário:", err)
			continue
		}
	}

	log.Println("User and projects updated successfully")
	return nil
}
