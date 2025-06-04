package repository

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model/converter"
	ticketModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket"
	ticketEntity "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket/repository/entity"
)

func (tr *ticketRepository) FindTicketById(id string) (ticketModel.TicketDomainInterface, *rest_err.RestErr) {
	ticketId, err := strconv.Atoi(id)
	if err != nil {
		return nil, rest_err.NewBadRequestError("Invalid ticket ID")
	}

	query := `SELECT id, title, request_user, sector, description, request_type, priority, attachment_urls, asana_task_id, status, project_id
			  FROM tickets WHERE id = ?`

	row := tr.databaseConnection.QueryRow(query, ticketId)

	var entity ticketEntity.TicketEntity
	err = row.Scan(
		&entity.ID,
		&entity.Title,
		&entity.RequestUser,
		&entity.Sector,
		&entity.Description,
		&entity.RequestType,
		&entity.Priority,
		&entity.AttachmentURLs,
		&entity.AsanaTaskID,
		&entity.Status,
		&entity.ProjectID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, rest_err.NewNotFoundError("Ticket not found")
		}
		return nil, rest_err.NewInternalServerError(fmt.Sprintf("Error fetching ticket: %s", err.Error()))
	}

	asanaTaskId := entity.AsanaTaskID

	log.Println("Asana Task ID:", asanaTaskId)

	domain := converter.ConvertTicketEntityToDomain(entity)

	comments, RestErr := tr.FindCommentsByTicketID(id)

	if RestErr != nil {
		log.Println("Error fetching comments:", RestErr)
		return nil, rest_err.NewInternalServerError("Error fetching comments")
	}
	domain.SetAsanaTaskID(asanaTaskId)
	domain.SetComments(comments)

	return domain, nil
}

func (tr *ticketRepository) FindAllTicketsByEmail(email string) ([]ticketModel.TicketDomainInterface, *rest_err.RestErr) {
	query := `SELECT id, title, request_user, sector, description, request_type, priority, attachment_urls, asana_task_id, status, project_id
			  FROM tickets WHERE request_user = ?`

	rows, err := tr.databaseConnection.Query(query, email)
	if err != nil {
		log.Println("Error fetching tickets by email:", err)
		return nil, rest_err.NewInternalServerError("Error fetching tickets")
	}
	defer rows.Close()

	tickets := make([]ticketModel.TicketDomainInterface, 0)

	for rows.Next() {
		var entity ticketEntity.TicketEntity
		if err := rows.Scan(
			&entity.ID,
			&entity.Title,
			&entity.RequestUser,
			&entity.Sector,
			&entity.Description,
			&entity.RequestType,
			&entity.Priority,
			&entity.AttachmentURLs,
			&entity.AsanaTaskID,
			&entity.Status,
			&entity.ProjectID,
		); err != nil {
			log.Println("Error scanning ticket:", err)
			continue
		}

		domain := converter.ConvertTicketEntityToDomain(entity)

		ticketIDStr := strconv.Itoa(int(entity.ID))
		comments, errRest := tr.FindCommentsByTicketID(ticketIDStr)
		if errRest == nil {
			domain.SetAsanaTaskID(entity.AsanaTaskID)
			domain.SetComments(comments)
		}

		tickets = append(tickets, domain)
	}

	return tickets, nil
}

func (tr *ticketRepository) FindAllTickets() ([]ticketModel.TicketDomainInterface, *rest_err.RestErr) {
	query := `SELECT id, title, request_user, sector, description, request_type, priority, attachment_urls, asana_task_id, status, project_id FROM tickets`

	rows, err := tr.databaseConnection.Query(query)
	if err != nil {
		log.Println("Error fetching all tickets:", err)
		return nil, rest_err.NewInternalServerError("Error fetching tickets")
	}
	defer rows.Close()

	tickets := make([]ticketModel.TicketDomainInterface, 0)

	for rows.Next() {
		var entity ticketEntity.TicketEntity
		if err := rows.Scan(
			&entity.ID,
			&entity.Title,
			&entity.RequestUser,
			&entity.Sector,
			&entity.Description,
			&entity.RequestType,
			&entity.Priority,
			&entity.AttachmentURLs,
			&entity.AsanaTaskID,
			&entity.Status,
			&entity.ProjectID,
		); err != nil {
			log.Println("Error scanning ticket:", err)
			continue
		}

		domain := converter.ConvertTicketEntityToDomain(entity)

		ticketIDStr := strconv.Itoa(int(entity.ID))
		comments, errRest := tr.FindCommentsByTicketID(ticketIDStr)
		if errRest == nil {
			domain.SetAsanaTaskID(entity.AsanaTaskID)
			domain.SetComments(comments)
		}

		tickets = append(tickets, domain)
	}

	return tickets, nil
}

func (tr *ticketRepository) FindCommentsByTicketID(ticketId string) ([]ticketModel.CommentDomain, *rest_err.RestErr) {
	var ticketIDInt int
	_, err := fmt.Sscanf(ticketId, "%d", &ticketIDInt)
	if err != nil {
		return nil, rest_err.NewBadRequestError("Invalid ticket ID")
	}

	rows, err := tr.databaseConnection.Query(`
		SELECT id, ticket_id, author, content, created_at
		FROM comments
		WHERE ticket_id = ?
		ORDER BY created_at ASC`, ticketIDInt)

	if err != nil {
		log.Println("Error fetching comments:", err)
		return nil, rest_err.NewInternalServerError("Error fetching comments")
	}
	defer rows.Close()

	var comments []ticketModel.CommentDomain
	for rows.Next() {
		var c ticketModel.CommentDomain
		var createdAt time.Time

		if err := rows.Scan(&c.ID, &c.TicketID, &c.Author, &c.Content, &createdAt); err != nil {
			log.Println("Error reading comment:", err)
			continue
		}

		c.CreatedAt = createdAt.Unix() // convert to int64 (Unix timestamp)
		comments = append(comments, c)
	}

	return comments, nil
}

func (tr *ticketRepository) FindCommentsByEmail(email string) ([]ticketModel.CommentDomain, *rest_err.RestErr) {
	rows, err := tr.databaseConnection.Query(`
		SELECT id, ticket_id, author, content, created_at
		FROM comments
		WHERE author = ?
		ORDER BY created_at ASC`, email)

	if err != nil {
		log.Println("Error fetching comments by email:", err)
		return nil, rest_err.NewInternalServerError("Error fetching comments")
	}
	defer rows.Close()

	var comments []ticketModel.CommentDomain
	for rows.Next() {
		var c ticketModel.CommentDomain
		var createdAt time.Time

		if err := rows.Scan(&c.ID, &c.TicketID, &c.Author, &c.Content, &createdAt); err != nil {
			log.Println("Error reading comment:", err)
			continue
		}

		c.CreatedAt = createdAt.Unix()
		comments = append(comments, c)
	}

	return comments, nil
}
