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
		}
	}

	return ticket, nil
}

func (td *ticketDomainService) FindTicketByEmailServices(email string) (ticketModel.TicketDomainInterface, *rest_err.RestErr) {

	return td.ticketRepository.FindTicketByEmail(email)
}
