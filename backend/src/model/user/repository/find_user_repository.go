package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model/converter"
	userModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user"
	userEntity "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user/repository/entity"
)

// FindUserByEmail finds a user by email from the database
func (ur *userRepository) FindUserByEmail(email string) (userModel.UserDomainInterface, *rest_err.RestErr) {
	log.Println("Querying user by email:", email)

	query := `
		SELECT id, name, email, password, phone, enterprise, department, role
		FROM users WHERE email = ?
	`

	row := ur.db.QueryRow(query, email)

	var entity userEntity.UserEntity

	err := row.Scan(
		&entity.ID,
		&entity.Name,
		&entity.Email,
		&entity.Password,
		&entity.Phone,
		&entity.Enterprise,
		&entity.Department,
		&entity.Role,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("User not found with email:", email)
			return nil, rest_err.NewNotFoundError(fmt.Sprintf("User with email %s not found", email))
		}
		log.Println("Error scanning user by email:", err)
		return nil, rest_err.NewInternalServerError(fmt.Sprintf("Error finding user: %s", err.Error()))
	}

	projects, restErr := ur.getUserProjects(entity.ID)
	if restErr != nil {
		return nil, restErr
	}

	log.Println("User found with email:", email)
	return converter.ConvertUserEntityToDomain(entity, projects), nil
}

// FindUserById finds a user by ID from the database
func (ur *userRepository) FindUserById(id string) (userModel.UserDomainInterface, *rest_err.RestErr) {
	log.Println("Querying user by ID:", id)

	query := `
		SELECT id, name, email, password, phone, enterprise, department, role
		FROM users WHERE id = ?
	`

	row := ur.db.QueryRow(query, id)

	var entity userEntity.UserEntity

	err := row.Scan(
		&entity.ID,
		&entity.Name,
		&entity.Email,
		&entity.Password,
		&entity.Phone,
		&entity.Enterprise,
		&entity.Department,
		&entity.Role,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("User not found with ID:", id)
			return nil, rest_err.NewNotFoundError(fmt.Sprintf("User with id %s not found", id))
		}
		log.Println("Error scanning user by ID:", err)
		return nil, rest_err.NewInternalServerError(fmt.Sprintf("Error finding user by id: %s", err.Error()))
	}

	projects, restErr := ur.getUserProjects(entity.ID)
	if restErr != nil {
		return nil, restErr
	}

	log.Println("User found with ID:", id)
	return converter.ConvertUserEntityToDomain(entity, projects), nil
}

// FindUserByEmailAndPassword searches for a user using both email and password
func (ur *userRepository) FindUserByEmailAndPassword(email, password string) (userModel.UserDomainInterface, *rest_err.RestErr) {
	log.Println("Querying user by email and password")

	query := `
		SELECT id, name, email, password, phone, enterprise, department, role
		FROM users WHERE email = ? AND password = ?
	`

	row := ur.db.QueryRow(query, email, password)

	var entity userEntity.UserEntity

	err := row.Scan(
		&entity.ID,
		&entity.Name,
		&entity.Email,
		&entity.Password,
		&entity.Phone,
		&entity.Enterprise,
		&entity.Department,
		&entity.Role,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Invalid credentials for user:", email)
			return nil, rest_err.NewForbiddenError("User or password is invalid")
		}
		log.Println("Error scanning user by email and password:", err)
		return nil, rest_err.NewInternalServerError(fmt.Sprintf("Error finding user: %s", err.Error()))
	}

	projects, restErr := ur.getUserProjects(entity.ID)
	if restErr != nil {
		return nil, restErr
	}

	log.Println("User authenticated successfully:", email)
	return converter.ConvertUserEntityToDomain(entity, projects), nil
}

// getUserProjects retrieves the list of project names associated with a given user ID
func (ur *userRepository) getUserProjects(userID int) ([]string, *rest_err.RestErr) {
	log.Println("Fetching projects for user ID:", userID)

	query := `
		SELECT p.name FROM projects p
		INNER JOIN user_projects up ON up.project_id = p.id
		WHERE up.user_id = ?
	`

	rows, err := ur.db.Query(query, userID)
	if err != nil {
		log.Println("Error querying user projects:", err)
		return nil, rest_err.NewInternalServerError("Error retrieving user's projects")
	}
	defer rows.Close()

	var projects []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err == nil {
			projects = append(projects, name)
		} else {
			log.Println("Error scanning project name:", err)
		}
	}

	log.Printf("Found %d projects for user ID %d\n", len(projects), userID)
	return projects, nil
}
