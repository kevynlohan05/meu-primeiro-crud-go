package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/controller"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface, ticketController controller.TicketControllerInterface) {

	r.GET("/user/getUserById/:userId", userController.FindUserById)
	r.GET("/user/getUserByEmail/:userEmail", userController.FindUserByEmail)
	r.POST("/user/createUser", userController.CreateUser)
	r.PUT("/user/updateUser/:userId", userController.UpdateUser)
	r.DELETE("/user/deleteUser/:userId", userController.DeleteUser)

	r.POST("/ticket/createTicket", ticketController.CreateTicket)
	r.GET("/ticket/getTicketById/:ticketId", ticketController.FindTicketById)
	r.PUT("/ticket/updateTicket/:ticketId", ticketController.UpdateTicket)
	r.DELETE("/ticket/deleteTicket/:ticketId", ticketController.DeleteTicket)
}
