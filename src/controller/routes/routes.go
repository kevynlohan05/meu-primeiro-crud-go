package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserById/:userId", controller.FindUserById)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)
	r.POST("/createTicket", controller.CreateTicket)
	r.GET("/getTicketById/:ticketId", controller.FindTicketById)
	r.PUT("/updateTicket/:ticketId", controller.UpdateTicket)
	r.DELETE("/deleteTicket/:ticketId", controller.DeleteTicket)
}
