package converter

import (
	"log"

	ticketModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket"
	ticketEntity "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket/repository/entity"
	userModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user"
	userEntity "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user/repository/entity"
)

func ConvertUserEntityToDomain(entity userEntity.UserEntity) userModel.UserDomainInterface {
	// Criando o domínio a partir da entidade
	domain := userModel.NewUserDomain(
		entity.Name,
		entity.Email,
		entity.Password,
		entity.Department,
		entity.Role,
	)

	domain.SetID(entity.ID.Hex())

	// Logando o domínio após a definição do ID
	log.Println("Domain after setting ID:")
	log.Printf("Domain: %+v\n", domain)

	return domain
}

func ConvertTicketEntityToDomain(entity ticketEntity.TicketEntity) ticketModel.TicketDomainInterface {
	// Criando o domínio a partir da entidade
	domain := ticketModel.NewTicketDomain(
		entity.Title,
		entity.RequestUser,
		entity.Sector,
		entity.Description,
		entity.RequestType,
		entity.Priority,
		entity.AttachmentURL,
	)

	domain.SetID(entity.ID.Hex())
	domain.SetAsanaTaskID(entity.AsanaTaskID)

	// Logando o domínio após a definição do ID
	log.Println("Domain after setting ID:")
	log.Printf("Domain: %+v\n", domain)

	return domain
}
