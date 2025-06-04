package view

import (
	"log"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/controller/model/response"
	projectModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/projects"
	ticketModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket"
	userModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user"
)

// ConvertUserDomainToResponse converts the User domain model into a UserResponse DTO for API response
func ConvertUserDomainToResponse(userDomain userModel.UserDomainInterface) response.UserResponse {
	projectsDomain := userDomain.GetProjects()

	// Create a copy of the projects slice to avoid mutation issues
	projectsResponse := make([]string, len(projectsDomain))
	copy(projectsResponse, projectsDomain)

	log.Printf("Converting user ID %s with %d projects to response", userDomain.GetID(), len(projectsDomain))

	return response.UserResponse{
		ID:         userDomain.GetID(),
		Name:       userDomain.GetName(),
		Email:      userDomain.GetEmail(),
		Phone:      userDomain.GetPhone(),
		Department: userDomain.GetDepartment(),
		Projects:   projectsResponse,
		Enterprise: userDomain.GetEnterprise(),
		Role:       userDomain.GetRole(),
	}
}

// ConvertTicketDomainToResponse converts the Ticket domain model to TicketResponse DTO including nested comments
func ConvertTicketDomainToResponse(ticketDomain ticketModel.TicketDomainInterface) response.TicketResponse {
	commentsDomain := ticketDomain.GetComments()
	commentsResponse := make([]response.CommentResponse, len(commentsDomain))

	for i, comment := range commentsDomain {
		commentsResponse[i] = response.CommentResponse{
			ID:        comment.ID,
			TicketID:  comment.TicketID,
			Author:    comment.Author,
			Content:   comment.Content,
			CreatedAt: comment.CreatedAt,
		}
	}

	log.Printf("Converting ticket ID %s with %d comments to response", ticketDomain.GetID(), len(commentsDomain))

	return response.TicketResponse{
		ID:             ticketDomain.GetID(),
		Status:         ticketDomain.GetStatus(),
		Title:          ticketDomain.GetTitle(),
		RequestUser:    ticketDomain.GetRequestUser(),
		Sector:         ticketDomain.GetSector(),
		Description:    ticketDomain.GetDescription(),
		RequestType:    ticketDomain.GetRequestType(),
		Priority:       ticketDomain.GetPriority(),
		AttachmentURLs: ticketDomain.GetAttachmentURLs(),
		AsanaTaskID:    ticketDomain.GetAsanaTaskID(),
		Comments:       commentsResponse,
		ProjectName:    ticketDomain.GetProjectName(),
	}
}

// ConvertProjectDomainToResponse converts the Project domain model into ProjectResponse DTO
func ConvertProjectDomainToResponse(projectDomain projectModel.ProjectDomainInterface) response.ProjectResponse {
	log.Printf("Converting project ID %s with name '%s' to response", projectDomain.GetID(), projectDomain.GetName())

	return response.ProjectResponse{
		ID:   projectDomain.GetID(),
		Name: projectDomain.GetName(),
	}
}
