package service

import (
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	ticketModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket"
)

func (td *ticketDomainService) FindTicketByIdServices(id string) (ticketModel.TicketDomainInterface, *rest_err.RestErr) {

	return td.ticketRepository.FindTicketById(id)
}

func (td *ticketDomainService) FindTicketByEmailServices(email string) (ticketModel.TicketDomainInterface, *rest_err.RestErr) {

	return td.ticketRepository.FindTicketByEmail(email)
}
