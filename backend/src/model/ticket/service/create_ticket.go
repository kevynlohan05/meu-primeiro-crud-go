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

	projectName := ticketDomain.GetProjectName()

	// Fetch user by email to verify permissions
	user, err := ud.userService.FindUserByEmailServices(ticketDomain.GetRequestUser())
	if err != nil {
		log.Println("Error fetching user by email:", err)
		return nil, err
	}

	// Check if user has access to the project
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

	// Retrieve project details by name
	project, err := ud.projectService.FindProjectByNameServices(projectName)
	if err != nil {
		log.Println("Project not found:", err)
		return nil, rest_err.NewBadRequestError("Invalid project name")
	}
	ticketDomain.SetAsanaProjectID(project.GetIdAsana())

	// Convert project ID string to int64
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

	// Create ticket in repository
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

	// Async: create task in Asana and upload attachments
	go func(ticket ticketModel.TicketDomainInterface) {
		log.Println("Starting Asana task creation")

		projectIDStr := strconv.FormatInt(ticket.GetProjectID(), 10)
		project, err := ud.projectService.FindProjectByIdServices(projectIDStr)
		if err != nil {
			log.Println("Error fetching project to get AsanaProjectID:", err)
			return
		}

		ticket.SetAsanaProjectID(project.GetIdAsana())

		taskID, RestErr := integrationAsana.CreateAsanaTask(ticket)
		if RestErr != nil {
			log.Println("Error creating task in Asana:", err)
			return
		}
		log.Printf("Asana task successfully created! ID: %s\n", taskID)

		if restErr := ud.ticketRepository.UpdateAsanaTaskID(ticket.GetID(), taskID); restErr != nil {
			log.Println("Error updating ticket with Asana task ID:", restErr)
			return
		}
		log.Println("Ticket updated with Asana task ID")

		// Upload attachments to Asana
		for _, filePath := range ticket.GetAttachmentURLs() {
			log.Println("Uploading file to Asana:", filePath)
			RestErr = integrationAsana.UploadAttachmentToAsana(taskID, filePath)
			if RestErr != nil {
				log.Println("Error uploading attachment to Asana:", err)
			} else {
				log.Println("File uploaded successfully to Asana:", filePath)

				// Remove local file after upload
				if removeErr := os.Remove(filePath); removeErr != nil {
					log.Println("Error deleting local file:", removeErr)
				} else {
					log.Println("Local file removed:", filePath)
				}
			}
		}
	}(ticketDomainRepository)

	// Set project name in the returned domain object
	ticketDomainRepository.SetProjectName(projectName)

	// Set initial status of ticket
	ticketDomainRepository.SetStatus("Novas")

	return ticketDomainRepository, nil
}
