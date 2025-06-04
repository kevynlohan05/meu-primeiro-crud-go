package repository

import (
	"fmt"
	"log"
	"strconv"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model/converter"
	ticketModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket"
)

func (tr *ticketRepository) UpdateTicket(ticketId string, ticketDomain ticketModel.TicketDomainInterface) *rest_err.RestErr {
	ticketIDInt, err := strconv.Atoi(ticketId)
	if err != nil {
		return rest_err.NewBadRequestError("Invalid ticket ID")
	}

	value := converter.ConvertTicketDomainToEntity(ticketDomain)
	if value == nil {
		return rest_err.NewInternalServerError("Failed to convert domain to entity")
	}

	query := `UPDATE tickets SET 
		title = ?, 
		request_user = ?, 
		sector = ?, 
		description = ?, 
		request_type = ?, 
		priority = ?, 
		attachment_urls = ?, 
		asana_task_id = ?, 
		status = ?, 
		project_id = ?, 
		comments = ?
		WHERE id = ?`

	result, err := tr.databaseConnection.Exec(query,
		value.Title,
		value.RequestUser,
		value.Sector,
		value.Description,
		value.RequestType,
		value.Priority,
		value.AttachmentURLs,
		value.AsanaTaskID,
		value.Status,
		value.ProjectID,
		value.Comments,
		ticketIDInt,
	)

	if err != nil {
		log.Println("Error updating ticket:", err)
		return rest_err.NewInternalServerError(fmt.Sprintf("Error updating ticket: %s", err.Error()))
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error getting affected rows count:", err)
		return rest_err.NewInternalServerError("Internal error updating ticket")
	}

	if rowsAffected == 0 {
		return rest_err.NewNotFoundError("Ticket not found for update")
	}

	log.Println("Ticket updated successfully in MySQL")
	return nil
}

func (tr *ticketRepository) UpdateAsanaTaskID(ticketId string, taskID string) *rest_err.RestErr {
	ticketIDInt, err := strconv.Atoi(ticketId)
	if err != nil {
		log.Println("Error converting ticket ID:", err)
		return rest_err.NewBadRequestError("Invalid ticket ID")
	}

	query := `UPDATE tickets SET asana_task_id = ? WHERE id = ?`

	result, err := tr.databaseConnection.Exec(query, taskID, ticketIDInt)
	if err != nil {
		log.Println("Error updating Asana task ID:", err)
		return rest_err.NewInternalServerError("Error updating Asana task ID")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error getting affected rows count:", err)
		return rest_err.NewInternalServerError("Internal error updating Asana task ID")
	}

	if rowsAffected == 0 {
		return rest_err.NewNotFoundError("Ticket not found to update Asana task ID")
	}

	log.Println("Asana task ID updated successfully in MySQL")
	return nil
}

func (tr *ticketRepository) UpdateTicketStatus(ticketId string, status string) *rest_err.RestErr {
	ticketIDInt, err := strconv.Atoi(ticketId)
	if err != nil {
		log.Println("Error converting ticket ID:", err)
		return rest_err.NewBadRequestError("Invalid ticket ID")
	}

	query := `UPDATE tickets SET status = ? WHERE id = ?`

	result, err := tr.databaseConnection.Exec(query, status, ticketIDInt)
	if err != nil {
		log.Println("Error updating ticket status:", err)
		return rest_err.NewInternalServerError("Error updating ticket status")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error getting affected rows count:", err)
		return rest_err.NewInternalServerError("Internal error updating ticket status")
	}

	if rowsAffected == 0 {
		return rest_err.NewNotFoundError("Ticket not found to update status")
	}

	log.Println("Ticket status updated successfully in MySQL")
	return nil
}
