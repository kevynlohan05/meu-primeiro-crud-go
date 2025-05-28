package repository

import (
	"database/sql"
	"encoding/json"
	"log"
	"fmt"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	ticketModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket"
)

func (tr *ticketRepository) AddComment(ticketId string, comment ticketModel.CommentDomain) *rest_err.RestErr {
	// Converter ticketId para int
	var ticketIDInt int
	_, err := fmt.Sscanf(ticketId, "%d", &ticketIDInt)
	if err != nil {
		return rest_err.NewBadRequestError("ID do ticket inválido")
	}

	// 1. Buscar campo comments atual no ticket
	var commentsJSON sql.NullString
	err = tr.databaseConnection.QueryRow("SELECT comments FROM tickets WHERE id = ?", ticketIDInt).Scan(&commentsJSON)
	if err != nil {
		log.Println("Erro ao buscar comentários do ticket:", err)
		return rest_err.NewInternalServerError("Erro ao buscar ticket")
	}

	// 2. Desserializar JSON para slice
	var comments []ticketModel.CommentDomain
	if commentsJSON.Valid && commentsJSON.String != "" {
		err = json.Unmarshal([]byte(commentsJSON.String), &comments)
		if err != nil {
			log.Println("Erro ao desserializar comentários:", err)
			return rest_err.NewInternalServerError("Erro ao processar comentários")
		}
	}

	// 3. Adicionar novo comentário
	comments = append(comments, comment)

	// 4. Serializar novamente para JSON
	newCommentsJSON, err := json.Marshal(comments)
	if err != nil {
		log.Println("Erro ao serializar comentários:", err)
		return rest_err.NewInternalServerError("Erro ao processar comentários")
	}

	// 5. Atualizar o campo comments no banco
	_, err = tr.databaseConnection.Exec("UPDATE tickets SET comments = ? WHERE id = ?", string(newCommentsJSON), ticketIDInt)
	if err != nil {
		log.Println("Erro ao atualizar comentários no ticket:", err)
		return rest_err.NewInternalServerError("Erro ao salvar comentário")
	}

	return nil
}
