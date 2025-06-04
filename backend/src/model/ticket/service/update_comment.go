package service

import (
	"fmt"
	"log"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	ticketModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket"
)

func (td *ticketDomainService) UpdateComment(ticketId, commentId, email string, comment ticketModel.CommentDomain) *rest_err.RestErr {
	log.Println("Calling repository to update comment")

	var ticketIDInt int64
	_, err := fmt.Sscanf(ticketId, "%d", &ticketIDInt)
	if err != nil {
		log.Println("Invalid ticket ID format:", err)
		return rest_err.NewBadRequestError("Invalid ticket ID format")
	}

	var commentIDInt int64
	_, err = fmt.Sscanf(commentId, "%d", &commentIDInt)
	if err != nil {
		log.Println("Invalid comment ID format:", err)
		return rest_err.NewBadRequestError("Invalid comment ID format")
	}

	commentsUser, RestErr := td.ticketRepository.FindCommentsByEmail(email)
	if RestErr != nil {
		log.Println("Error retrieving comments by user:", RestErr)
		return rest_err.NewInternalServerError("Error retrieving comments")
	}

	var commentFound bool
	for _, c := range commentsUser {
		if c.ID == commentIDInt && c.TicketID == ticketIDInt {
			commentFound = true
			break
		}
	}
	if !commentFound {
		log.Println("Comment not found or user not authorized to update it")
		return rest_err.NewNotFoundError("Comment not found or user not authorized to update it")
	}

	RestErr = td.ticketRepository.UpdateComment(ticketId, commentId, comment.GetContent())
	if RestErr != nil {
		log.Println("Repository error while updating comment:", RestErr)
		return RestErr
	}

	log.Println("Comment updated successfully")
	return nil
}
