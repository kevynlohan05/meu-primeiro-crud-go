package service

import (
	"log"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	ticketModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket"
)

func (ud *ticketDomainService) CreateTicket(ticketDomain ticketModel.TicketDomainInterface) (ticketModel.TicketDomainInterface, *rest_err.RestErr) {

	log.Println("Calling repository to create ticket")
	ticketDomainRepository, err := ud.ticketRepository.CreateTicket(ticketDomain)
	if err != nil {
		log.Println("Error in repository:", err)
		return nil, err
	}

	if ticketDomainRepository == nil {
		log.Println("Error: ticketDomainRepository is nil")
		return nil, rest_err.NewInternalServerError("Failed to create ticket in repository")
	}

	log.Println("Ticket created successfully")

	return ticketDomainRepository, nil
}
