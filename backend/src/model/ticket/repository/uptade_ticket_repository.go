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
		return rest_err.NewBadRequestError("ID do ticket inválido")
	}

	value := converter.ConvertTicketDomainToEntity(ticketDomain)
	if value == nil {
		return rest_err.NewInternalServerError("Falha na conversão do domínio para entidade")
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
		log.Println("Erro ao atualizar ticket:", err)
		return rest_err.NewInternalServerError(fmt.Sprintf("Erro ao atualizar ticket: %s", err.Error()))
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Erro ao obter número de linhas afetadas:", err)
		return rest_err.NewInternalServerError("Erro interno ao atualizar ticket")
	}

	if rowsAffected == 0 {
		return rest_err.NewNotFoundError("Ticket não encontrado para atualização")
	}

	log.Println("Ticket atualizado com sucesso no MySQL")
	return nil
}

func (tr *ticketRepository) UpdateAsanaTaskID(ticketId string, taskID string) *rest_err.RestErr {
	ticketIDInt, err := strconv.Atoi(ticketId)
	if err != nil {
		log.Println("Erro ao converter ID do ticket:", err)
		return rest_err.NewBadRequestError("ID do ticket inválido")
	}

	query := `UPDATE tickets SET asana_task_id = ? WHERE id = ?`

	result, err := tr.databaseConnection.Exec(query, taskID, ticketIDInt)
	if err != nil {
		log.Println("Erro ao atualizar taskID do Asana:", err)
		return rest_err.NewInternalServerError("Erro ao atualizar taskID do Asana")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Erro ao obter número de linhas afetadas:", err)
		return rest_err.NewInternalServerError("Erro interno ao atualizar taskID do Asana")
	}

	if rowsAffected == 0 {
		return rest_err.NewNotFoundError("Ticket não encontrado para atualizar Asana Task ID")
	}

	log.Println("Task ID do Asana atualizado com sucesso no MySQL")
	return nil
}
