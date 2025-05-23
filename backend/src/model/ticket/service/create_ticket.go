package service

import (
	"log"
	"os"

	integrationAsana "github.com/kevynlohan05/meu-primeiro-crud-go/src/integration"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	ticketModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket"
)

func (ud *ticketDomainService) CreateTicket(ticketDomain ticketModel.TicketDomainInterface) (ticketModel.TicketDomainInterface, *rest_err.RestErr) {
	ticketDomain.SetStatus("Novas")
	log.Println("Starting ticket creation process")

	// Fetch user by email
	user, err := ud.userService.FindUserByEmailServices(ticketDomain.GetRequestUser())
	if err != nil {
		log.Println("Error fetching user by email:", err)
		return nil, err
	}

	// Check if user has access to the selected project
	hasAccess := false
	for _, project := range user.GetProjects() {
		if project == ticketDomain.GetProjects() {
			hasAccess = true
			break
		}
	}

	if !hasAccess {
		log.Println("User is not allowed to create a ticket in this project")
		return nil, rest_err.NewBadRequestError("User is not allowed to create a ticket in this project")
	}

	// Create the ticket in the repository
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

	// Launch asynchronous process to create the Asana task
	go func(ticket ticketModel.TicketDomainInterface) {
		log.Println("Starting Asana task creation")
		taskID, err := integrationAsana.CreateAsanaTask(ticket)
		if err != nil {
			log.Println("Error while creating Asana task:", err)
			return
		}

		log.Printf("Asana task successfully created! ID: %s\n", taskID)

		// Update ticket with Asana task ID
		if restErr := ud.ticketRepository.UpdateAsanaTaskID(ticket.GetID(), taskID); restErr != nil {
			log.Println("Error updating ticket with Asana task ID:", restErr)
			return
		}
		log.Println("Ticket updated with Asana task ID")

		// 🔽 Upload dos arquivos anexados
		for _, filePath := range ticket.GetAttachmentURLs() {
			log.Println("Uploading file to Asana:", filePath)

			err = integrationAsana.UploadAttachmentToAsana(taskID, filePath)
			if err != nil {
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

	return ticketDomainRepository, nil
}
