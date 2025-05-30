package service

import (
	"log"
	"os"
	"strconv"

	integrationAsana "github.com/kevynlohan05/meu-primeiro-crud-go/src/integration"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	ticketModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket"
)

func (ud *ticketDomainService) CreateTicket(ticketDomain ticketModel.TicketDomainInterface) (ticketModel.TicketDomainInterface, *rest_err.RestErr) {
	log.Println("Starting ticket creation process")
	ticketDomain.SetStatus("Novas")

	projectName := ticketDomain.GetProjectName()

	user, err := ud.userService.FindUserByEmailServices(ticketDomain.GetRequestUser())
	if err != nil {
		log.Println("Error fetching user by email:", err)
		return nil, err
	}

	hasAccess := false
	for _, userProject := range user.GetProjects() {
		if userProject == projectName {
			hasAccess = true
			break
		}
	}
	if !hasAccess {
		log.Println("User is not allowed to create a ticket in this project")
		return nil, rest_err.NewBadRequestError("User is not allowed to create a ticket in this project")
	}

	project, err := ud.projectService.FindProjectByNameServices(projectName)
	if err != nil {
		log.Println("Project not found:", err)
		return nil, rest_err.NewBadRequestError("Invalid project name")
	}
	ticketDomain.SetAsanaProjectID(project.GetIdAsana())

	projectIDStr := project.GetID()
	var projectIDInt int64
	if parsedID, err := strconv.ParseInt(projectIDStr, 10, 64); err == nil {
		projectIDInt = parsedID
	} else {
		log.Println("Error converting project ID to int64:", err)
		return nil, rest_err.NewInternalServerError("Failed to convert project ID")
	}
	ticketDomain.SetProjectID(projectIDInt)
	log.Println(ticketDomain.GetAsanaProjectID(), "is the Asana project ID")

	ticketDomainRepository, err := ud.ticketRepository.CreateTicket(ticketDomain)
	if err != nil {
		log.Println("Repository error while creating ticket:", err)
		return nil, err
	}
	if ticketDomainRepository == nil {
		log.Println("Repository returned nil while creating ticket")
		return nil, rest_err.NewInternalServerError("Failed to create ticket in repository")
	}

	log.Println("Ticket successfully created")

	go func(ticket ticketModel.TicketDomainInterface) {
		log.Println("Starting Asana task creation")

		projectIDStr := strconv.FormatInt(ticket.GetProjectID(), 10)
		project, err := ud.projectService.FindProjectByIdServices(projectIDStr)
		if err != nil {
			log.Println("Erro ao buscar projeto para obter AsanaProjectID:", err)
			return
		}

		ticket.SetAsanaProjectID(project.GetIdAsana())

		taskID, RestErr := integrationAsana.CreateAsanaTask(ticket)
		if RestErr != nil {
			log.Println("Erro ao criar tarefa no Asana:", err)
			return
		}
		log.Printf("Asana task successfully created! ID: %s\n", taskID)

		if restErr := ud.ticketRepository.UpdateAsanaTaskID(ticket.GetID(), taskID); restErr != nil {
			log.Println("Erro ao atualizar ticket com ID da tarefa Asana:", restErr)
			return
		}
		log.Println("Ticket atualizado com ID da tarefa Asana")

		for _, filePath := range ticket.GetAttachmentURLs() {
			log.Println("Uploading file to Asana:", filePath)
			RestErr = integrationAsana.UploadAttachmentToAsana(taskID, filePath)
			if RestErr != nil {
				log.Println("Erro ao enviar anexo para o Asana:", err)
			} else {
				log.Println("Arquivo enviado com sucesso ao Asana:", filePath)

				if removeErr := os.Remove(filePath); removeErr != nil {
					log.Println("Erro ao apagar arquivo local:", removeErr)
				} else {
					log.Println("Arquivo local removido:", filePath)
				}
			}
		}
	}(ticketDomainRepository)

	ticketDomainRepository.SetProjectName(projectName)

	return ticketDomainRepository, nil
}
