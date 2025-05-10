package service

import (
	"log"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
)

func (td *ticketDomainService) DeleteTicket(ticketId string) *rest_err.RestErr {
	log.Println("Calling repository to delete ticket")

	err := td.ticketRepository.DeleteTicket(ticketId)
	if err != nil {
		log.Println("Error in repository:", err)
		return err
	}

	log.Println("Ticket delete successfully")

	return nil
}
