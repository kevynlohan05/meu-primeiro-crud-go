package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/controller/model/request"
	ticketModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket"
)

func (tc *ticketControllerInterface) UpdateComment(c *gin.Context) {
	ticketId := c.Param("ticketId")
	commentId := c.Param("commentId")

	if ticketId == "" || commentId == "" {
		errorMessage := rest_err.NewBadRequestError("Ticket ID and Comment ID are required")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	var commentUpdateRequest request.UpdateCommentRequest
	if err := c.ShouldBindJSON(&commentUpdateRequest); err != nil {
		errorMessage := rest_err.NewBadRequestError("Invalid request body")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userEmail := c.GetString("userEmail")
	if userEmail == "" {
		errorMessage := rest_err.NewUnauthorizedError("User not authenticated")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	comment := ticketModel.CommentDomain{
		Content: commentUpdateRequest.Content,
	}

	err := tc.service.UpdateComment(ticketId, commentId, userEmail, comment)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment updated successfully"})
}
