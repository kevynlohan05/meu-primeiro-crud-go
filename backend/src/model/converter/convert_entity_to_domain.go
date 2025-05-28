package converter

import (
	"encoding/json"
	"log"
	"fmt"

	ticketModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket"
	ticketEntity "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket/repository/entity"
	userModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user"
	userEntity "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user/repository/entity"
)

func ConvertUserEntityToDomain(entity userEntity.UserEntity) userModel.UserDomainInterface {
	var projects []string
	if err := json.Unmarshal([]byte(entity.Projects), &projects); err != nil {
		log.Printf("Error unmarshaling projects JSON: %v\n", err)
		projects = []string{} 
	}

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

	log.Println("Domain after setting ID:")
	log.Printf("Domain: %+v\n", domain)

	return domain
}

func ConvertTicketEntityToDomain(entity ticketEntity.TicketEntity) ticketModel.TicketDomainInterface {
	var attachmentURLs []string
	err := json.Unmarshal([]byte(entity.AttachmentURLs), &attachmentURLs)
	if err != nil {
		log.Printf("Error unmarshaling attachment URLs: %v\n", err)
		attachmentURLs = []string{} // fallback vazio
	}


	domain := ticketModel.NewTicketDomain(
		entity.Title,
		entity.RequestUser,
		entity.Sector,
		entity.Description,
		entity.RequestType,
		entity.Priority,
		entity.Projects,
		attachmentURLs,
	)

	domain.SetID(fmt.Sprintf("%d", entity.ID))
	domain.SetAsanaTaskID(entity.AsanaTaskID)

	var comments []ticketModel.CommentDomain
	err = json.Unmarshal([]byte(entity.Comments), &comments)
	if err != nil {
		log.Printf("Error unmarshaling comments: %v\n", err)
		comments = []ticketModel.CommentDomain{}
	}
	domain.SetComments(comments)

	// Logando o domínio após a definição do ID
	log.Println("Domain after setting ID:")
	log.Printf("Domain: %+v\n", domain)

	return domain
}
