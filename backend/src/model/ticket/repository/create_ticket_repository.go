package repository

import (
	"fmt"
	"log"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model/converter"
	ticketModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket"
)

func (tr *ticketRepository) CreateTicket(ticketDomain ticketModel.TicketDomainInterface) (ticketModel.TicketDomainInterface, *rest_err.RestErr) {
	log.Println("Converting domain to entity")

	value := converter.ConvertTicketDomainToEntity(ticketDomain)

	if value == nil {
		log.Println("Error: Conversion to entity failed")
		return nil, rest_err.NewInternalServerError("Failed to convert domain to entity")
	}

	query := `
		INSERT INTO tickets 
		(title, request_user, sector, description, request_type, priority, attachment_urls, asana_task_id, status, project_id)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	result, err := tr.databaseConnection.Exec(
		query,
		value.Title,
		value.RequestUser,
		value.Sector,
		value.Description,
		value.RequestType,
		value.Priority,
		value.AttachmentURLs,
		value.AsanaTaskID,
		value.Status,
		value.ProjectID, // adicionado corretamente
	)

	if err != nil {
		log.Println("Error inserting ticket into MySQL:", err)
		return nil, rest_err.NewInternalServerError(fmt.Sprintf("Error inserting ticket: %v", err))
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		log.Println("Error getting inserted ticket ID:", err)
		return nil, rest_err.NewInternalServerError("Error retrieving ticket ID")
	}

	log.Println("Ticket inserted successfully into MySQL")

	value.ID = int64(insertedID)

	log.Println("Converting entity back to domain")
	return converter.ConvertTicketEntityToDomain(*value), nil
}
