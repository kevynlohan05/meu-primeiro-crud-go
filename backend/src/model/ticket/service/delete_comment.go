package service

import (
	"fmt"
	"log"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
)

func (td *ticketDomainService) DeleteComment(ticketId, commentId, email string) *rest_err.RestErr {
	log.Println("Calling repository to delete comment")

	// Validate the ticket ID format
	var ticketIDInt int64
	_, err := fmt.Sscanf(ticketId, "%d", &ticketIDInt)
	if err != nil {
		log.Println("Invalid ticket ID format:", err)
		return rest_err.NewBadRequestError("Invalid ticket ID format")
	}

	// Validate the comment ID format
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
	for _, comment := range commentsUser {
		if comment.ID == commentIDInt && comment.TicketID == ticketIDInt {
			commentFound = true
			break
		}
	}
	if !commentFound {
		log.Println("Comment not found or user not authorized to delete it")
		return rest_err.NewNotFoundError("Comment not found or user not authorized to delete it")
	}

	RestErr = td.ticketRepository.DeleteComment(ticketId, commentId)
	if RestErr != nil {
		log.Println("Repository error during comment deletion:", RestErr)
		return RestErr
	}

	log.Println("Comment deleted successfully")
	return nil
}
