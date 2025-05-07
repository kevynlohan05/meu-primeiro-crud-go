package converter

import (
	"log"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model/repository/entity"
)

func ConvertUserEntityToDomain(entity entity.UserEntity) model.UserDomainInterface {
	// Criando o domínio a partir da entidade
	domain := model.NewUserDomain(
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

func ConvertTicketEntityToDomain(entity entity.TicketEntity) model.TicketDomainInterface {
	// Criando o domínio a partir da entidade
	domain := model.NewTicketDomain(
		entity.Title,
		entity.Description,
		entity.RequestType,
		entity.Priority,
		entity.AttachmentURL,
	)

	domain.SetID(entity.ID.Hex())

	// Logando o domínio após a definição do ID
	log.Println("Domain after setting ID:")
	log.Printf("Domain: %+v\n", domain)

	return domain
}
