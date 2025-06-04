package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/controller/model/request"
	ticketModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket"
)

func (tc *ticketControllerInterface) AddComment(c *gin.Context) {
	log.Println("Init AddComment controller")

	var req request.AddCommentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid data for comment")
		c.JSON(errRest.Code, errRest)
		return
	}

	ticketId := c.Param("ticketId")
	if ticketId == "" {
		errRest := rest_err.NewBadRequestError("Ticket ID is required")
		c.JSON(errRest.Code, errRest)
		return
	}

	userEmail := c.GetString("userEmail")
	if userEmail == "" {
		errRest := rest_err.NewUnauthorizedError("User not authenticated")
		c.JSON(errRest.Code, errRest)
		return
	}

	comment := ticketModel.CommentDomain{
		Author:    userEmail,
		Content:   req.Content,
		CreatedAt: time.Now().Unix(),
	}

	err := tc.service.AddComment(ticketId, comment)
	if err != nil {
		log.Println("Error adding comment:", err)
		c.JSON(err.Code, err)
		return
	}

	c.Status(http.StatusOK)
}
