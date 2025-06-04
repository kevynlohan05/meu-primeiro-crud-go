package repository

import (
	"fmt"
	"log"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
)

func (tr *ticketRepository) UpdateComment(ticketId string, commentId string, content string) *rest_err.RestErr {
	var ticketIDInt int
	var commentIDInt int

	_, err := fmt.Sscanf(ticketId, "%d", &ticketIDInt)
	if err != nil {
		return rest_err.NewBadRequestError("ID do ticket inválido")
	}

	_, err = fmt.Sscanf(commentId, "%d", &commentIDInt)
	if err != nil {
		return rest_err.NewBadRequestError("ID do comentário inválido")
	}

	_, err = tr.databaseConnection.Exec(`
		UPDATE comments
		SET content = ?
		WHERE ticket_id = ? AND id = ?`,
		content, ticketIDInt, commentIDInt)

	if err != nil {
		log.Println("Erro ao atualizar comentário:", err)
		return rest_err.NewInternalServerError("Erro ao atualizar comentário")
	}

	return nil
}
