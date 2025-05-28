package repository

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model/converter"
	ticketModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket"
	ticketEntity "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket/repository/entity"
)

func (tr *ticketRepository) FindTicketById(id string) (ticketModel.TicketDomainInterface, *rest_err.RestErr) {
	// Converte string para int
	ticketId, err := strconv.Atoi(id)
	if err != nil {
		return nil, rest_err.NewBadRequestError("ID do ticket inválido")
	}

	// Query para buscar ticket pelo ID
	query := `SELECT id, title, request_user, department, description, request_type, priority, attachment_url, asana_task_id, status FROM tickets WHERE id = ?`

	row := tr.databaseConnection.QueryRow(query, ticketId)

	var entity ticketEntity.TicketEntity
	err = row.Scan(
		&entity.ID,
		&entity.Title,
		&entity.RequestUser,
		&entity.Department,
		&entity.Description,
		&entity.RequestType,
		&entity.Priority,
		&entity.AttachmentURLs,
		&entity.AsanaTaskID,
		&entity.Status,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, rest_err.NewNotFoundError("Ticket não encontrado")
		}
		return nil, rest_err.NewInternalServerError(fmt.Sprintf("Erro ao buscar ticket: %s", err.Error()))
	}

	return converter.ConvertTicketEntityToDomain(entity), nil
}

func (tr *ticketRepository) FindAllTicketsByEmail(email string) ([]ticketModel.TicketDomainInterface, *rest_err.RestErr) {
	query := `SELECT id, title, request_user, department, description, request_type, priority, attachment_url, asana_task_id, status 
			  FROM tickets WHERE request_user = ?`

	rows, err := tr.databaseConnection.Query(query, email)
	if err != nil {
		log.Println("Erro ao buscar tickets por email:", err)
		return nil, rest_err.NewInternalServerError("Erro ao buscar tickets")
	}
	defer rows.Close()

	tickets := make([]ticketModel.TicketDomainInterface, 0)

	for rows.Next() {
		var entity ticketEntity.TicketEntity
		err := rows.Scan(
			&entity.ID,
			&entity.Title,
			&entity.RequestUser,
			&entity.Department,
			&entity.Description,
			&entity.RequestType,
			&entity.Priority,
			&entity.AttachmentURLs,
			&entity.AsanaTaskID,
			&entity.Status,
		)
		if err != nil {
			log.Println("Erro ao escanear ticket:", err)
			return nil, rest_err.NewInternalServerError("Erro ao processar tickets")
		}
		tickets = append(tickets, converter.ConvertTicketEntityToDomain(entity))
	}

	if len(tickets) == 0 {
		return nil, rest_err.NewNotFoundError("Nenhum ticket encontrado para o e-mail informado")
	}

	return tickets, nil
}

func (tr *ticketRepository) FindAllTickets() ([]ticketModel.TicketDomainInterface, *rest_err.RestErr) {
	query := `SELECT id, title, request_user, department, description, request_type, priority, attachment_url, asana_task_id, status FROM tickets`

	rows, err := tr.databaseConnection.Query(query)
	if err != nil {
		log.Println("Erro ao buscar todos os tickets:", err)
		return nil, rest_err.NewInternalServerError("Erro ao buscar tickets")
	}
	defer rows.Close()

	tickets := make([]ticketModel.TicketDomainInterface, 0)

	for rows.Next() {
		var entity ticketEntity.TicketEntity
		err := rows.Scan(
			&entity.ID,
			&entity.Title,
			&entity.RequestUser,
			&entity.Department,
			&entity.Description,
			&entity.RequestType,
			&entity.Priority,
			&entity.AttachmentURLs,
			&entity.AsanaTaskID,
			&entity.Status,
		)
		if err != nil {
			log.Println("Erro ao escanear ticket:", err)
			return nil, rest_err.NewInternalServerError("Erro ao processar tickets")
		}
		tickets = append(tickets, converter.ConvertTicketEntityToDomain(entity))
	}

	if len(tickets) == 0 {
		return nil, rest_err.NewNotFoundError("Nenhum ticket encontrado")
	}

	return tickets, nil
}
