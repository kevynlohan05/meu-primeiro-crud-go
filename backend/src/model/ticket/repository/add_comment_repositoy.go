package repository

import (
	"fmt"
	"log"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	ticketModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket"
)

func (tr *ticketRepository) AddComment(ticketId string, comment ticketModel.CommentDomain) *rest_err.RestErr {
	var ticketIDInt int
	_, err := fmt.Sscanf(ticketId, "%d", &ticketIDInt)
	if err != nil {
		return rest_err.NewBadRequestError("Invalid ticket ID")
	}

	_, err = tr.databaseConnection.Exec(`
		INSERT INTO comments (ticket_id, author, content)
		VALUES (?, ?, ?)`,
		ticketIDInt, comment.Author, comment.Content)

	if err != nil {
		log.Println("Error adding comment:", err)
		return rest_err.NewInternalServerError("Error saving comment")
	}

	return nil
}
