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
	log.Println("Start CreateTicket controller")

	// Parse form fields
	ticketRequest := request.TicketRequest{
		Title:       c.PostForm("title"),
		Description: c.PostForm("description"),
		RequestType: c.PostForm("request_type"),
		Priority:    c.PostForm("priority"),
		Project:     c.PostForm("project"),
	}

	// Validate request fields
	if err := validate.Struct(ticketRequest); err != nil {
		log.Println("Validation error:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid input data",
			"details": err.Error(),
		})
		return
	}

	// Retrieve user info from context (from middleware)
	userEmail := c.GetString("userEmail")
	userDepartment := c.GetString("userDepartment")

	// Parse and save uploaded files
	form, err := c.MultipartForm()
	if err != nil {
		log.Println("Failed to read multipart form:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to process uploaded files"})
		return
	}

	files := form.File["attachment_urls"]
	var uploadedURLs []string

	for _, file := range files {
		path := "uploads/" + file.Filename
		if err := c.SaveUploadedFile(file, path); err != nil {
			log.Println("Failed to save file:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}
		uploadedURLs = append(uploadedURLs, path)
	}

	// Build domain object
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

	// Call service to create the ticket
	domainResult, restErr := tc.service.CreateTicket(domain)
	if restErr != nil {
		log.Println("Failed to create ticket:", restErr)
		c.JSON(restErr.Code, restErr)
		return
	}

	// Success response
	c.JSON(http.StatusOK, view.ConvertTicketDomainToResponse(domainResult))
}
