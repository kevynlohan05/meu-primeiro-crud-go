package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/validation"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/controller/model/request"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/controller/model/response"
)

func CreateTicket(c *gin.Context) {

	log.Println("Init createTicket controller")
	var ticketRequest request.TicketRequest

	if err := c.ShouldBindJSON(&ticketRequest); err != nil {
		errRest := validation.ValidateRequestError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(ticketRequest)
	response := response.TicketResponse{
		Title:       ticketRequest.Title,
		Description: ticketRequest.Description,
		Priority:    ticketRequest.Priority,
		RequestType: ticketRequest.RequestType,
	}

	c.JSON(http.StatusOK, response)
}
