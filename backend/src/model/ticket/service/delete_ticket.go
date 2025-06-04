package service

import (
	"log"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
)

func (td *ticketDomainService) DeleteTicket(ticketId string) *rest_err.RestErr {
	log.Println("Calling repository to delete ticket with ID:", ticketId)

	err := td.ticketRepository.DeleteTicket(ticketId)
	if err != nil {
		log.Println("Repository error while deleting ticket:", err)
		return err
	}

	log.Println("Ticket deleted successfully with ID:", ticketId)

	return nil
}
