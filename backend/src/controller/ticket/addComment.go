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
		errRest := rest_err.NewBadRequestError("Dados inválidos para comentário")
		c.JSON(errRest.Code, errRest)
		return
	}

	ticketId := c.Param("ticketId")
	if ticketId == "" {
		errRest := rest_err.NewBadRequestError("ID do ticket é obrigatório")
		c.JSON(errRest.Code, errRest)
		return
	}

	userEmail := c.GetString("userEmail")

	comment := ticketModel.CommentDomain{
		Author:    userEmail,
		Message:   req.Message,
		Timestamp: time.Now().Unix(),
	}

	err := tc.service.AddComment(ticketId, comment)
	if err != nil {
		log.Println("Error adding comment:", err)
		c.JSON(err.Code, err)
		return
	}

	c.Status(http.StatusOK)
}
