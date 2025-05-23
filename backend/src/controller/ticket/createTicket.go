package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/controller/model/request"
	ticketModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/view"
)

var validate = validator.New()

func (tc *ticketControllerInterface) CreateTicket(c *gin.Context) {
	log.Println("Init createTicket controller")

	// Preenche manualmente
	ticketRequest := request.TicketRequest{
		Title:       c.PostForm("title"),
		Description: c.PostForm("description"),
		RequestType: c.PostForm("request_type"),
		Priority:    c.PostForm("priority"),
		Project:     c.PostForm("project"),
	}

	// Valida com o validator
	if err := validate.Struct(ticketRequest); err != nil {
		log.Println("Erro de validação:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	userEmail := c.GetString("userEmail")
	userDepartment := c.GetString("userDepartment")

	form, err := c.MultipartForm()
	if err != nil {
		log.Println("Erro ao ler MultipartForm:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Falha ao processar arquivos"})
		return
	}

	files := form.File["attachment_urls"]
	var uploadedURLs []string

	for _, file := range files {
		path := "uploads/" + file.Filename
		if err := c.SaveUploadedFile(file, path); err != nil {
			log.Println("Erro ao salvar arquivo:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao salvar arquivo"})
			return
		}
		uploadedURLs = append(uploadedURLs, path)
	}

	domain := ticketModel.NewTicketDomain(
		ticketRequest.Title,
		userEmail,
		userDepartment,
		ticketRequest.Description,
		ticketRequest.RequestType,
		ticketRequest.Priority,
		ticketRequest.Project,
		uploadedURLs,
	)

	domainResult, RestErr := tc.service.CreateTicket(domain)
	if RestErr != nil {
		log.Println("Erro ao criar ticket:", RestErr)
		c.JSON(RestErr.Code, RestErr)
		return
	}

	c.JSON(http.StatusOK, view.ConvertTicketDomainToResponse(domainResult))
}
