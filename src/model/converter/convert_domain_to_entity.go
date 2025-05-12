package converter

import (
	"log"

	ticketModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket"
	ticketEntity "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket/repository/entity"
	userModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user"
	userEntity "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user/repository/entity"
)

func ConvertUserDomainToEntity(domain userModel.UserDomainInterface) *userEntity.UserEntity {

	// Criando a entidade
	entity := &userEntity.UserEntity{
		Name:       domain.GetName(),
		Email:      domain.GetEmail(),
		Password:   domain.GetPassword(),
		Department: domain.GetDepartment(),
		Role:       domain.GetRole(),
	}

	// Logando o que foi convertido para a entidade
	log.Println("Converted domain to entity successfully:")
	log.Printf("Entity: %+v\n", entity)

	return entity
}

func ConvertTicketDomainToEntity(domain ticketModel.TicketDomainInterface) *ticketEntity.TicketEntity {
	// Criando a entidade
	entity := &ticketEntity.TicketEntity{
		Title:         domain.GetTitle(),
		Description:   domain.GetDescription(),
		RequestType:   domain.GetRequestType(),
		Priority:      domain.GetPriority(),
		AttachmentURL: domain.GetAttachmentURL(),
		UserEmail:     domain.GetUserEmail(),
	}

	// Logando o que foi convertido para a entidade
	log.Println("Converted domain to entity successfully:")
	log.Printf("Entity: %+v\n", entity)

	return entity
}
