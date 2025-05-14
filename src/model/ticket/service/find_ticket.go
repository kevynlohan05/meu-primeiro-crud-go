package service

import (
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	integrationAsana "github.com/kevynlohan05/meu-primeiro-crud-go/src/integration"
	ticketModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket"
)

func (td *ticketDomainService) FindTicketByIdServices(id string) (ticketModel.TicketDomainInterface, *rest_err.RestErr) {
	ticket, err := td.ticketRepository.FindTicketById(id)
	if err != nil {
		return nil, err
	}

	if ticket.GetAsanaTaskID() != "" {
		status, _, err := integrationAsana.GetAsanaTaskDetails(ticket.GetAsanaTaskID())
		if err == nil {
			ticket.SetStatus(status)
			td.ticketRepository.UpdateTicket(id, ticket)
		}
	}

	return ticket, nil
}

func (ts *ticketDomainService) FindAllTicketsByUser(email string) ([]ticketModel.TicketDomainInterface, *rest_err.RestErr) {
	tickets, err := ts.ticketRepository.FindAllTicketsByEmail(email)
	if err != nil {
		return nil, err
	}

	for _, t := range tickets {
		asanaTaskID := t.GetAsanaTaskID()
		if asanaTaskID != "" {
			status, _, err := integrationAsana.GetAsanaTaskDetails(asanaTaskID)
			if err == nil {
				t.SetStatus(status)
				_ = ts.ticketRepository.UpdateTicket(t.GetID(), t) // atualiza no banco
			}
		}
	}

	return tickets, nil
}

func (ts *ticketDomainService) FindAllTickets() ([]ticketModel.TicketDomainInterface, *rest_err.RestErr) {
	tickets, err := ts.ticketRepository.FindAllTickets()
	if err != nil {
		return nil, err
	}

	for _, t := range tickets {
		asanaTaskID := t.GetAsanaTaskID()
		if asanaTaskID != "" {
			status, _, err := integrationAsana.GetAsanaTaskDetails(asanaTaskID)
			if err == nil {
				t.SetStatus(status)
				_ = ts.ticketRepository.UpdateTicket(t.GetID(), t) // atualiza no banco
			}
		}
	}

	return tickets, nil
}
