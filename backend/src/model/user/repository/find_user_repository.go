package repository

import (
	"database/sql"
	"fmt"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model/converter"
	userModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user"
	userEntity "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user/repository/entity"
)

func (ur *userRepository) FindUserByEmail(email string) (userModel.UserDomainInterface, *rest_err.RestErr) {
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
			return nil, rest_err.NewNotFoundError(fmt.Sprintf("User with email %s not found", email))
		}
		return nil, rest_err.NewInternalServerError(fmt.Sprintf("Error finding user: %s", err.Error()))
	}

	projects, restErr := ur.getUserProjects(entity.ID)
	if restErr != nil {
		return nil, restErr
	}

	return converter.ConvertUserEntityToDomain(entity, projects), nil
}

func (ur *userRepository) FindUserById(id string) (userModel.UserDomainInterface, *rest_err.RestErr) {
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
			return nil, rest_err.NewNotFoundError(fmt.Sprintf("User with id %s not found", id))
		}
		return nil, rest_err.NewInternalServerError(fmt.Sprintf("Error finding user by id: %s", err.Error()))
	}

	projects, restErr := ur.getUserProjects(entity.ID)
	if restErr != nil {
		return nil, restErr
	}

	return converter.ConvertUserEntityToDomain(entity, projects), nil
}

func (ur *userRepository) FindUserByEmailAndPassword(email, password string) (userModel.UserDomainInterface, *rest_err.RestErr) {
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
			return nil, rest_err.NewForbiddenError("User or password is invalid")
		}
		return nil, rest_err.NewInternalServerError(fmt.Sprintf("Error finding user: %s", err.Error()))
	}

	projects, restErr := ur.getUserProjects(entity.ID)
	if restErr != nil {
		return nil, restErr
	}

	return converter.ConvertUserEntityToDomain(entity, projects), nil
}

func (ur *userRepository) getUserProjects(userID int) ([]string, *rest_err.RestErr) {
	query := `
		SELECT p.name FROM projects p
		INNER JOIN user_projects up ON up.project_id = p.id
		WHERE up.user_id = ?
	`

	rows, err := ur.db.Query(query, userID)
	if err != nil {
		return nil, rest_err.NewInternalServerError("Erro ao buscar projetos do usuário")
	}
	defer rows.Close()

	var projects []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err == nil {
			projects = append(projects, name)
		}
	}

	return projects, nil
}
