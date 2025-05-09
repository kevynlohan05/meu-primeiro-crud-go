package service

import (
	"log"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model"
)

func (td *ticketDomainService) UpdateTicket(ticketId string, ticketDomain model.TicketDomainInterface) *rest_err.RestErr {
	log.Println("Calling repository to update ticket")

	err := td.ticketRepository.UpdateTicket(ticketId, ticketDomain)
	if err != nil {
		log.Println("Error in repository:", err)
		return err
	}

	log.Println("Ticket update successfully")

	return nil
}
