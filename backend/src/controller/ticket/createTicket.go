package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/validation"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/controller/model/request"
	ticketModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/view"
)

var (
	TicketDomainInterface ticketModel.TicketDomainInterface
)

func (tc *ticketControllerInterface) CreateTicket(c *gin.Context) {

	log.Println("Init createTicket controller")
	var ticketRequest request.TicketRequest

	if err := c.ShouldBindJSON(&ticketRequest); err != nil {
		errRest := validation.ValidateRequestError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	userEmail := c.GetString("userEmail")

	userDepartment := c.GetString("userDepartment")

	form, err := c.MultipartForm()
	if err != nil {
		log.Println("Error parsing multipart form:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid multipart form"})
		return
	}

	files := form.File["attachments"]
	var uploadedURLs []string

	for _, file := range files {
		// Caminho onde o arquivo será salvo localmente
		path := "uploads/" + file.Filename

		if err := c.SaveUploadedFile(file, path); err != nil {
			log.Println("Error saving uploaded file:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
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

	domainResult, err := tc.service.CreateTicket(domain)
	if err != nil {
		log.Println("Error creating ticket:", err)
		errRest := validation.ValidateRequestError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, view.ConvertTicketDomainToResponse(domainResult))
}
