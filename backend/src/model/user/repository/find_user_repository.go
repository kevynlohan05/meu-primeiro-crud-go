package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model/converter"
	userModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user"
	userEntity "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user/repository/entity"
)

func (ur *userRepository) FindUserByEmail(email string) (userModel.UserDomainInterface, *rest_err.RestErr) {
	query := `
		SELECT id, name, email, password, phone, enterprise, department, role, projects
		FROM users WHERE email = ?
	`

	row := ur.db.QueryRow(query, email)

	var entity userEntity.UserEntity
	var projectsJSON string

	err := row.Scan(
		&entity.ID,
		&entity.Name,
		&entity.Email,
		&entity.Password,
		&entity.Phone,
		&entity.Enterprise,
		&entity.Department,
		&entity.Role,
		&projectsJSON,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, rest_err.NewNotFoundError(fmt.Sprintf("User with email %s not found", email))
		}
		return nil, rest_err.NewInternalServerError(fmt.Sprintf("Error finding user: %s", err.Error()))
	}

	if err := json.Unmarshal([]byte(projectsJSON), &entity.Projects); err != nil {
		return nil, rest_err.NewInternalServerError("Failed to parse user projects")
	}

	return converter.ConvertUserEntityToDomain(entity), nil
}

func (ur *userRepository) FindUserById(id string) (userModel.UserDomainInterface, *rest_err.RestErr) {
	query := `
		SELECT id, name, email, password, phone, enterprise, department, role, projects
		FROM users WHERE id = ?
	`

	row := ur.db.QueryRow(query, id)

	var entity userEntity.UserEntity
	var projectsJSON string

	err := row.Scan(
		&entity.ID,
		&entity.Name,
		&entity.Email,
		&entity.Password,
		&entity.Phone,
		&entity.Enterprise,
		&entity.Department,
		&entity.Role,
		&projectsJSON,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, rest_err.NewNotFoundError(fmt.Sprintf("User with id %s not found", id))
		}
		return nil, rest_err.NewInternalServerError(fmt.Sprintf("Error finding user by id: %s", err.Error()))
	}

	if err := json.Unmarshal([]byte(projectsJSON), &entity.Projects); err != nil {
		return nil, rest_err.NewInternalServerError("Failed to parse user projects")
	}

	return converter.ConvertUserEntityToDomain(entity), nil
}

func (ur *userRepository) FindUserByEmailAndPassword(email, password string) (userModel.UserDomainInterface, *rest_err.RestErr) {
	query := `
		SELECT id, name, email, password, phone, enterprise, department, role, projects
		FROM users WHERE email = ? AND password = ?
	`

	row := ur.db.QueryRow(query, email, password)

	var entity userEntity.UserEntity
	var projectsJSON string

	err := row.Scan(
		&entity.ID,
		&entity.Name,
		&entity.Email,
		&entity.Password,
		&entity.Phone,
		&entity.Enterprise,
		&entity.Department,
		&entity.Role,
		&projectsJSON,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, rest_err.NewForbiddenError("User or password is invalid")
		}
		return nil, rest_err.NewInternalServerError(fmt.Sprintf("Error finding user: %s", err.Error()))
	}

	if err := json.Unmarshal([]byte(projectsJSON), &entity.Projects); err != nil {
		return nil, rest_err.NewInternalServerError("Failed to parse user projects")
	}

	return converter.ConvertUserEntityToDomain(entity), nil
}
