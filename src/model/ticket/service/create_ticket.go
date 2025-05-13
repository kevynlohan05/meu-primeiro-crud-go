package service

import (
	"log"

	integrationAsana "github.com/kevynlohan05/meu-primeiro-crud-go/src/integration"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	ticketModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket"
)

func (ud *ticketDomainService) CreateTicket(ticketDomain ticketModel.TicketDomainInterface) (ticketModel.TicketDomainInterface, *rest_err.RestErr) {

	ticketDomain.SetStatus("Novas")
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

	go func(ticket ticketModel.TicketDomainInterface) {
		log.Println("Chamando integração com o Asana")
		taskID, err := integrationAsana.CreateAsanaTask(ticket)
		if err != nil {
			log.Println("Erro ao criar tarefa no Asana:", err)
			return
		}
		log.Printf("Tarefa criada com sucesso no Asana! ID: %s\n", taskID)

		// Atualizar o ticket com o ID da task do Asana
		restErr := ud.ticketRepository.UpdateAsanaTaskID(ticket.GetID(), taskID)
		if restErr != nil {
			log.Println("Erro ao atualizar ticket com taskID do Asana:", err)
			return
		}
		log.Println("ticket atualizado com taskID do Asana")
	}(ticketDomainRepository)

	return ticketDomainRepository, nil
}
