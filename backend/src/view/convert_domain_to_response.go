package view

import (
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/controller/model/response"
	ticketModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket"
	userModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user"
)

func ConvertUserDomainToResponse(userDomain userModel.UserDomainInterface) response.UserResponse {
	projectsDomain := userDomain.GetProjects()
	projectsResponse := make([]string, len(projectsDomain))

	for i, project := range projectsDomain {
		projectsResponse[i] = project
	}

	return response.UserResponse{
		ID:         userDomain.GetID(),
		Name:       userDomain.GetName(),
		Email:      userDomain.GetEmail(),
		Department: userDomain.GetDepartment(),
		Projects:   projectsResponse,
		Role:       userDomain.GetRole(),
	}
}

func ConvertTicketDomainToResponse(ticketDomain ticketModel.TicketDomainInterface) response.TicketResponse {
	commentsDomain := ticketDomain.GetComments()
	commentsResponse := make([]response.CommentResponse, len(commentsDomain))

	for i, comment := range commentsDomain {
		commentsResponse[i] = response.CommentResponse{
			Author:    comment.Author,
			Message:   comment.Message,
			Timestamp: comment.Timestamp,
		}
	}

	return response.TicketResponse{
		ID:            ticketDomain.GetID(),
		Status:        ticketDomain.GetStatus(),
		Title:         ticketDomain.GetTitle(),
		RequestUser:   ticketDomain.GetRequestUser(),
		Sector:        ticketDomain.GetSector(),
		Description:   ticketDomain.GetDescription(),
		RequestType:   ticketDomain.GetRequestType(),
		Priority:      ticketDomain.GetPriority(),
		AttachmentURL: ticketDomain.GetAttachmentURL(),
		AsanaTaskID:   ticketDomain.GetAsanaTaskID(),
		Comments:      commentsResponse,
		Projects:      ticketDomain.GetProjects(),
	}
}
