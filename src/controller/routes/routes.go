package routes

import (
	"github.com/gin-gonic/gin"
	controllerTicket "github.com/kevynlohan05/meu-primeiro-crud-go/src/controller/ticket"
	controllerUser "github.com/kevynlohan05/meu-primeiro-crud-go/src/controller/user"
	userModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user"
)

func InitRoutes(r *gin.RouterGroup, userController controllerUser.UserControllerInterface, ticketController controllerTicket.TicketControllerInterface) {

	r.GET("/user/getUserById/:userId", userModel.VerifyTokenMiddleware, userController.FindUserById)
	r.GET("/user/getUserByEmail/:userEmail", userModel.VerifyTokenMiddleware, userController.FindUserByEmail)
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
