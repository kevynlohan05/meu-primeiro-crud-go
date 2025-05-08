package service

import (
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model"
)

func (td *ticketDomainService) FindTicketByIdServices(id string) (model.TicketDomainInterface, *rest_err.RestErr) {

	return td.ticketRepository.FindTicketById(id)
}

func (td *ticketDomainService) FindTicketByEmailServices(email string) (model.TicketDomainInterface, *rest_err.RestErr) {

	return td.ticketRepository.FindTicketByEmail(email)
}
