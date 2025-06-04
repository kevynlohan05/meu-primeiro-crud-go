package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/controller/model/request"
	ticketModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket"
)

func (tc *ticketControllerInterface) UpdateComment(c *gin.Context) {
	log.Println("Init UpdateComment controller")

	ticketId := c.Param("ticketId")
	commentId := c.Param("commentId")

	if ticketId == "" || commentId == "" {
		log.Println("Ticket ID and Comment ID are required")
		errorMessage := rest_err.NewBadRequestError("Ticket ID and Comment ID are required")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	var commentUpdateRequest request.UpdateCommentRequest
	if err := c.ShouldBindJSON(&commentUpdateRequest); err != nil {
		log.Println("Invalid request body:", err)
		errorMessage := rest_err.NewBadRequestError("Invalid request body")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userEmail := c.GetString("userEmail")
	if userEmail == "" {
		log.Println("User not authenticated")
		errorMessage := rest_err.NewUnauthorizedError("User not authenticated")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	comment := ticketModel.CommentDomain{
		Content: commentUpdateRequest.Content,
	}

	if err := tc.service.UpdateComment(ticketId, commentId, userEmail, comment); err != nil {
		log.Printf("Error updating comment: %v\n", err)
		c.JSON(err.Code, err)
		return
	}

	log.Printf("Comment %s for ticket %s updated successfully by user %s\n", commentId, ticketId, userEmail)
	c.JSON(http.StatusOK, gin.H{"message": "Comment updated successfully"})
}
