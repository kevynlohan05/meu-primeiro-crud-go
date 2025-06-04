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
		return rest_err.NewBadRequestError("Invalid ticket ID")
	}

	_, err = fmt.Sscanf(commentId, "%d", &commentIDInt)
	if err != nil {
		return rest_err.NewBadRequestError("Invalid comment ID")
	}

	_, err = tr.databaseConnection.Exec(`
		UPDATE comments
		SET content = ?
		WHERE ticket_id = ? AND id = ?`,
		content, ticketIDInt, commentIDInt)

	if err != nil {
		log.Println("Error updating comment:", err)
		return rest_err.NewInternalServerError("Error updating comment")
	}

	return nil
}
