package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/view"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (tc *ticketControllerInterface) FindTicketById(c *gin.Context) {
	ticketId := c.Param("ticketId")

	if _, err := primitive.ObjectIDFromHex(ticketId); err != nil {
		errorMessage := rest_err.NewBadRequestError("Invalid ticket ID format")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	ticketDomain, err := tc.service.FindTicketByIdServices(ticketId)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertTicketDomainToResponse(ticketDomain))
}

func (tc *ticketControllerInterface) FindTicketByEmail(c *gin.Context) {

	ticketEmail := c.Param("ticketEmail")

	if _, err := uuid.Parse(ticketEmail); err != nil {
		errorMessage := rest_err.NewBadRequestError("Invalid ticket Email format")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	ticketDomain, err := tc.service.FindTicketByEmailServices(ticketEmail)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertTicketDomainToResponse(ticketDomain))
}
