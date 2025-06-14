package converter

import (
	"encoding/json"
	"fmt"
	"log"

	projectModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/projects"
	projectEntity "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/projects/repository/entity"
	ticketModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket"
	ticketEntity "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket/repository/entity"
	userModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user"
	userEntity "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user/repository/entity"
)

func ConvertUserEntityToDomain(entity userEntity.UserEntity, projects []string) userModel.UserDomainInterface {
	domain := userModel.NewUserDomain(
		entity.Name,
		entity.Email,
		entity.Password,
		entity.Phone,
		entity.Enterprise,
		entity.Department,
		entity.Role,
		projects,
	)

	domain.SetID(fmt.Sprintf("%d", entity.ID))
	return domain
}

func ConvertTicketEntityToDomain(entity ticketEntity.TicketEntity) ticketModel.TicketDomainInterface {
	var attachmentURLs []string
	err := json.Unmarshal([]byte(entity.AttachmentURLs), &attachmentURLs)
	if err != nil {
		log.Printf("Error unmarshaling attachment URLs: %v\n", err)
		attachmentURLs = []string{} // fallback to empty
	}

	domain := ticketModel.NewTicketDomainFromEntity(
		entity.Title,
		entity.RequestUser,
		entity.Sector,
		entity.Description,
		entity.RequestType,
		entity.Priority,
		entity.AsanaTaskID,
		attachmentURLs,
		entity.ProjectID,
	)

	domain.SetID(fmt.Sprintf("%d", entity.ID))
	domain.SetAsanaTaskID(entity.AsanaTaskID)
	domain.SetStatus(entity.Status)

	log.Println("Domain after setting ID:")
	log.Printf("Domain: %+v\n", domain)

	return domain
}

func ConvertProjectEntityToDomain(entity projectEntity.ProjectEntity) projectModel.ProjectDomainInterface {
	domain := projectModel.NewProjectDomain(
		entity.Name,
		entity.IdAsana,
	)

	domain.SetID(fmt.Sprintf("%d", entity.ID))
	return domain
}
