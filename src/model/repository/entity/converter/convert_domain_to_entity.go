package converter

import (
	"log"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model/repository/entity"
)

func ConvertUserDomainToEntity(domain model.UserDomainInterface) *entity.UserEntity {

	// Criando a entidade
	entity := &entity.UserEntity{
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

func ConvertTicketDomainToEntity(domain model.TicketDomainInterface) *entity.TicketEntity {
	// Criando a entidade
	entity := &entity.TicketEntity{
		Title:         domain.GetTitle(),
		Description:   domain.GetDescription(),
		RequestType:   domain.GetRequestType(),
		Priority:      domain.GetPriority(),
		AttachmentURL: domain.GetAttachmentURL(),
	}

	// Logando o que foi convertido para a entidade
	log.Println("Converted domain to entity successfully:")
	log.Printf("Entity: %+v\n", entity)

	return entity
}
