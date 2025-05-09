package view

import (
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/controller/model/response"
	ticketModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket"
	userModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user"
)

func ConvertUserDomainToResponse(userDomain userModel.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:         userDomain.GetID(),
		Name:       userDomain.GetName(),
		Email:      userDomain.GetEmail(),
		Department: userDomain.GetDepartment(),
		Role:       userDomain.GetRole(),
	}
}

func ConvertTicketDomainToResponse(ticketDomain ticketModel.TicketDomainInterface) response.TicketResponse {
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
