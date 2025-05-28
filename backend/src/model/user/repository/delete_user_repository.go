package repository

import (
	"fmt"
	"log"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
)

func (ur *userRepository) DeleteUser(userId string) *rest_err.RestErr {
	log.Println("Deleting user in MySQL with id:", userId)

	query := `DELETE FROM users WHERE id = ?`

	result, err := ur.db.Exec(query, userId)
	if err != nil {
		log.Println("Error deleting user from MySQL:", err)
		return rest_err.NewInternalServerError(fmt.Sprintf("Error deleting user: %v", err))
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error getting affected rows:", err)
		return rest_err.NewInternalServerError("Error verifying delete operation")
	}

	if rowsAffected == 0 {
		log.Println("No user found with id:", userId)
		return rest_err.NewNotFoundError(fmt.Sprintf("User with id %s not found", userId))
	}

	log.Println("User deleted successfully from MySQL")
	return nil
}
