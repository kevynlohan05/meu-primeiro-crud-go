package repository

import (
	"fmt"
	"log"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model/converter"
	userModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user"
)

// UpdateUser updates the user record and their associated projects in the database
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

	log.Printf("Executing update for user ID %s\n", userId)
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

	log.Println("Deleting old user_projects relationships")
	_, err = ur.db.Exec("DELETE FROM user_projects WHERE user_id = ?", userId)
	if err != nil {
		log.Println("Error deleting old user_projects:", err)
		return rest_err.NewInternalServerError("Error updating user projects")
	}

	// Insert new project relationships
	for _, projectName := range userDomain.GetProjects() {
		var projectID int

		log.Printf("Checking existence of project '%s'\n", projectName)
		err := ur.db.QueryRow("SELECT id FROM projects WHERE name = ?", projectName).Scan(&projectID)
		if err != nil {
			// If project does not exist, create it
			log.Printf("Project '%s' not found, inserting new project\n", projectName)
			res, err := ur.db.Exec("INSERT INTO projects (name) VALUES (?)", projectName)
			if err != nil {
				log.Println("Error inserting new project:", err)
				continue
			}
			lastID, _ := res.LastInsertId()
			projectID = int(lastID)
		}

		// Link user with project
		_, err = ur.db.Exec("INSERT INTO user_projects (user_id, project_id) VALUES (?, ?)", userId, projectID)
		if err != nil {
			log.Println("Error linking project to user:", err)
			continue
		}
		log.Printf("Linked project '%s' (ID %d) with user ID %s\n", projectName, projectID, userId)
	}

	log.Println("User and projects updated successfully")
	return nil
}
