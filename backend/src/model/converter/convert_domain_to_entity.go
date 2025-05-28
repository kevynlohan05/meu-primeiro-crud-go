package converter

import (
	"encoding/json"
	"log"

	ticketModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket"
	ticketEntity "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket/repository/entity"
	userModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user"
	userEntity "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user/repository/entity"
)

func ConvertUserDomainToEntity(domain userModel.UserDomainInterface) *userEntity.UserEntity {
	projectsJSON, err := json.Marshal(domain.GetProjects())
	if err != nil {
		log.Printf("Error marshaling projects: %v\n", err)
		return nil
	}

	entity := &userEntity.UserEntity{
		Name:       domain.GetName(),
		Email:      domain.GetEmail(),
		Password:   domain.GetPassword(),
		Phone:      domain.GetPhone(),
		Enterprise: domain.GetEnterprise(),
		Department: domain.GetDepartment(),
		Projects:   string(projectsJSON),
		Role:       domain.GetRole(),
	}

	log.Println("Converted domain to entity successfully:")
	log.Printf("Entity: %+v\n", entity)

	return entity
}

func ConvertTicketDomainToEntity(domain ticketModel.TicketDomainInterface) *ticketEntity.TicketEntity {
	attachmentURLsJSON, err := json.Marshal(domain.GetAttachmentURLs())
	if err != nil {
		log.Printf("Error marshaling attachment URLs: %v\n", err)
		attachmentURLsJSON = []byte("[]") // fallback vazio
	}

	entity := &ticketEntity.TicketEntity{
		Title:          domain.GetTitle(),
		RequestUser:    domain.GetRequestUser(),
		Sector:         domain.GetSector(),
		Description:    domain.GetDescription(),
		RequestType:    domain.GetRequestType(),
		Priority:       domain.GetPriority(),
		AttachmentURLs: string(attachmentURLsJSON),
		Status:         domain.GetStatus(),
		Projects:       domain.GetProjects(),
	}

	commentsJSON, err := json.Marshal(domain.GetComments())
	if err != nil {
		log.Printf("Error marshaling comments: %v\n", err)
		commentsJSON = []byte("[]")
	}
	entity.Comments = string(commentsJSON)

	log.Println("Converted domain to entity successfully:")
	log.Printf("Entity: %+v\n", entity)

	return entity
}
