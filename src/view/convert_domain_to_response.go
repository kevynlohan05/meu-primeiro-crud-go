package view

import (
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/controller/model/response"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model"
)

func ConvertUserDomainToResponse(userDomain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:         userDomain.GetID(),
		Name:       userDomain.GetName(),
		Email:      userDomain.GetEmail(),
		Department: userDomain.GetDepartment(),
		Role:       userDomain.GetRole(),
	}
}

func ConvertTicketDomainToResponse(ticketDomain model.TicketDomainInterface) response.TicketResponse {
	return response.TicketResponse{
		ID:            ticketDomain.GetID(),
		Title:         ticketDomain.GetTitle(),
		Description:   ticketDomain.GetDescription(),
		Status:        "Em andamento",
		RequestType:   ticketDomain.GetRequestType(),
		Priority:      ticketDomain.GetPriority(),
		AttachmentURL: ticketDomain.GetAttachmentURL(),
	}
}
