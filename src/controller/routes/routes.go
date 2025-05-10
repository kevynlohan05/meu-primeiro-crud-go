package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/controller"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface, ticketController controller.TicketControllerInterface) {

	r.GET("/user/getUserById/:userId", model.VerifyTokenMiddleware, userController.FindUserById)
	r.GET("/user/getUserByEmail/:userEmail", model.VerifyTokenMiddleware, userController.FindUserByEmail)
	r.POST("/user/createUser", userController.CreateUser)
	r.PUT("/user/updateUser/:userId", userController.UpdateUser)
	r.DELETE("/user/deleteUser/:userId", userController.DeleteUser)

	r.POST("/user/login", userController.LoginUser)

	r.POST("/ticket/createTicket", ticketController.CreateTicket)
	r.GET("/ticket/getTicketById/:ticketId", ticketController.FindTicketById)
	r.GET("/ticket/getTicketByEmail/:ticketEmail", ticketController.FindTicketByEmail)
	r.PUT("/ticket/updateTicket/:ticketId", ticketController.UpdateTicket)
	r.DELETE("/ticket/deleteTicket/:ticketId", ticketController.DeleteTicket)
}
