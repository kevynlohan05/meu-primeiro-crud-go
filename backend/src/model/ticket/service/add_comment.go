package service

import (
	"log"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	ticketModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket"
)

func (td *ticketDomainService) AddComment(ticketId string, comment ticketModel.CommentDomain) *rest_err.RestErr {
	log.Println("Calling repository to add comment")

	err := td.ticketRepository.AddComment(ticketId, comment)
	if err != nil {
		log.Println("Error in repository:", err)
		return err
	}

	log.Println("Comment added successfully")
	return nil
}
