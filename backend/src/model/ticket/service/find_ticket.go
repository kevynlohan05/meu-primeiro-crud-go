package service

import (
	"log"
	"strconv"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	integrationAsana "github.com/kevynlohan05/meu-primeiro-crud-go/src/integration"
	ticketModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket"
)

func (td *ticketDomainService) FindTicketByIdServices(id string) (ticketModel.TicketDomainInterface, *rest_err.RestErr) {
	ticket, err := td.ticketRepository.FindTicketById(id)
	if err != nil {
		return nil, err
	}

	projectIDStr := strconv.FormatInt(ticket.GetProjectID(), 10)

	project, err := td.projectService.FindProjectByIdServices(projectIDStr)
	if err != nil {
		return nil, rest_err.NewInternalServerError("Erro ao buscar o projeto associado ao ticket")
	}

	ticket.SetProjectName(project.GetName())

	if ticket.GetAsanaTaskID() != "" {
		status, _, err := integrationAsana.GetAsanaTaskDetails(ticket.GetAsanaTaskID())
		if err == nil {
			ticket.SetStatus(status)
			td.ticketRepository.UpdateTicketStatus(id, status)
		}
	}

	return ticket, nil
}

func (td *ticketDomainService) FindAllTicketsByEmail(email string) ([]ticketModel.TicketDomainInterface, *rest_err.RestErr) {
	tickets, err := td.ticketRepository.FindAllTicketsByEmail(email)
	if err != nil {
		return nil, err
	}

	for _, t := range tickets {

		projectIDStr := strconv.FormatInt(t.GetProjectID(), 10)
		project, _ := td.projectService.FindProjectByIdServices(projectIDStr)
		t.SetProjectName(project.GetName())

		asanaTaskID := t.GetAsanaTaskID()
		if asanaTaskID != "" {
			status, _, err := integrationAsana.GetAsanaTaskDetails(asanaTaskID)
			if err == nil {
				t.SetStatus(status)
				_ = td.ticketRepository.UpdateTicketStatus(t.GetID(), status)
			}
		}
	}

	return tickets, nil
}

func (td *ticketDomainService) FindAllTickets() ([]ticketModel.TicketDomainInterface, *rest_err.RestErr) {
	tickets, err := td.ticketRepository.FindAllTickets()
	if err != nil {
		return nil, err
	}

	for _, t := range tickets {

		projectIDStr := strconv.FormatInt(t.GetProjectID(), 10)
		project, _ := td.projectService.FindProjectByIdServices(projectIDStr)
		t.SetProjectName(project.GetName())

		asanaTaskID := t.GetAsanaTaskID()
		log.Println("Asana Task ID:", asanaTaskID)
		if asanaTaskID != "" {
			status, _, err := integrationAsana.GetAsanaTaskDetails(asanaTaskID)
			log.Println(err)

			if err == nil {
				t.SetStatus(status)
				log.Println("Atualizando status do ticket:", t.GetID(), "para", status)
				_ = td.ticketRepository.UpdateTicketStatus(t.GetID(), status)
			}

		}
	}

	return tickets, nil
}

func (td *ticketDomainService) FindAllTicketsByEmailAndStatus(email string, status string) ([]ticketModel.TicketDomainInterface, *rest_err.RestErr) {
	tickets, err := td.ticketRepository.FindAllTicketsByEmail(email)
	if err != nil {
		return nil, err
	}

	var filteredTickets []ticketModel.TicketDomainInterface

	for _, t := range tickets {

		projectIDStr := strconv.FormatInt(t.GetProjectID(), 10)
		project, _ := td.projectService.FindProjectByIdServices(projectIDStr)
		t.SetProjectName(project.GetName())

		asanaTaskID := t.GetAsanaTaskID()
		if asanaTaskID != "" {
			newStatus, _, err := integrationAsana.GetAsanaTaskDetails(asanaTaskID)
			if err == nil {
				t.SetStatus(newStatus)
				_ = td.ticketRepository.UpdateTicketStatus(t.GetID(), status)
			}
		}

		if t.GetStatus() == status {
			filteredTickets = append(filteredTickets, t)
		}
	}

	if len(filteredTickets) == 0 {
		return nil, rest_err.NewNotFoundError("Nenhum ticket encontrado com os filtros fornecidos")
	}

	return filteredTickets, nil
}
