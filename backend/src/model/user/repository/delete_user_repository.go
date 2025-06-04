package repository

import (
	"fmt"
	"log"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
)

func (ur *userRepository) DeleteUser(userId string) *rest_err.RestErr {
	log.Printf("Attempting to delete user with ID: %s", userId)

	query := `DELETE FROM users WHERE id = ?`

	result, err := ur.db.Exec(query, userId)
	if err != nil {
		log.Printf("Error while deleting user from database: %v", err)
		return rest_err.NewInternalServerError(fmt.Sprintf("Error deleting user: %v", err))
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error retrieving affected rows after delete: %v", err)
		return rest_err.NewInternalServerError("Error verifying delete operation")
	}

	if rowsAffected == 0 {
		log.Printf("No user found with ID: %s", userId)
		return rest_err.NewNotFoundError(fmt.Sprintf("User with ID %s not found", userId))
	}

	log.Println("User successfully deleted")
	return nil
}
