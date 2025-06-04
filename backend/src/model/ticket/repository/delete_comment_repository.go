package repository

import (
	"fmt"
	"log"
	"strconv"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
)

func (tr *ticketRepository) DeleteComment(ticketId string, commentId string) *rest_err.RestErr {

	ticketID, err := strconv.Atoi(ticketId)
	if err != nil {
		return rest_err.NewBadRequestError("ID do ticket inválido")
	}

	commentID, err := strconv.Atoi(commentId)
	if err != nil {
		return rest_err.NewBadRequestError("ID do comentário inválido")
	}

	log.Println("Deleting comment from MySQL")

	result, err := tr.databaseConnection.Exec("DELETE FROM comments WHERE ticket_id = ? AND id = ?", ticketID, commentID)
	if err != nil {
		log.Println("Error deleting comment from MySQL:", err)
		return rest_err.NewInternalServerError(fmt.Sprintf("Erro ao deletar comentário: %s", err.Error()))
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error fetching rows affected:", err)
		return rest_err.NewInternalServerError("Erro ao processar remoção do comentário")
	}

	if rowsAffected == 0 {
		return rest_err.NewNotFoundError("Comentário não encontrado")
	}

	log.Println("Comment deleted successfully from MySQL")
	return nil
}
