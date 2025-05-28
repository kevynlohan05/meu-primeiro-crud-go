package repository

import (
	"fmt"
	"log"
	"strconv"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
)

func (tr *ticketRepository) DeleteTicket(ticketId string) *rest_err.RestErr {
	// Converte o ticketId string para int (supondo que o ID no MySQL seja INT)
	id, err := strconv.Atoi(ticketId)
	if err != nil {
		return rest_err.NewBadRequestError("ID do ticket inválido")
	}

	log.Println("Deleting ticket from MySQL")

	result, err := tr.databaseConnection.Exec("DELETE FROM tickets WHERE id = ?", id)
	if err != nil {
		log.Println("Error deleting ticket from MySQL:", err)
		return rest_err.NewInternalServerError(fmt.Sprintf("Erro ao deletar ticket: %s", err.Error()))
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error fetching rows affected:", err)
		return rest_err.NewInternalServerError("Erro ao processar remoção do ticket")
	}

	if rowsAffected == 0 {
		return rest_err.NewNotFoundError("Ticket não encontrado")
	}

	log.Println("Ticket deleted successfully from MySQL")
	return nil
}
