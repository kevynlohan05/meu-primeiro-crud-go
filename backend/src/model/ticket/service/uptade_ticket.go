package service

import (
	"log"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	ticketModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket"
)

func (td *ticketDomainService) UpdateTicket(ticketId string, ticketDomain ticketModel.TicketDomainInterface) *rest_err.RestErr {
	log.Println("Calling repository to update ticket")

	err := td.ticketRepository.UpdateTicket(ticketId, ticketDomain)
	if err != nil {
		log.Println("Repository error:", err)
		return err
	}

	log.Println("Ticket updated successfully")

	return nil
}

func (td *ticketDomainService) UpdateAsanaTaskID(ticketId string, taskID string) *rest_err.RestErr {
	log.Println("Calling repository to update Asana task ID for ticket")

	err := td.ticketRepository.UpdateAsanaTaskID(ticketId, taskID)
	if err != nil {
		log.Println("Repository error:", err)
		return err
	}

	log.Println("Ticket updated successfully")

	return nil
}
