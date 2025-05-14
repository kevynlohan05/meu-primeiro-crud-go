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
	r.PUT("/user/updateUser/:userId", userModel.VerifyTokenMiddleware, userController.UpdateUser)
	r.DELETE("/user/deleteUser/:userId", userModel.VerifyTokenMiddleware, userController.DeleteUser)

	r.POST("/user/login", userController.LoginUser)

	r.POST("/ticket/createTicket", userModel.VerifyTokenMiddleware, ticketController.CreateTicket)
	r.GET("/ticket/getTicketById/:ticketId", userModel.VerifyTokenMiddleware, ticketController.FindTicketById)
	r.GET("/ticket/getAllTicketsByEmail/:ticketEmail", userModel.VerifyTokenMiddleware, ticketController.FindAllTicketsByUser)
	r.GET("/ticket/getAllTickets", userModel.VerifyTokenMiddleware, ticketController.FindAllTickets)
	r.PUT("/ticket/updateTicket/:ticketId", userModel.VerifyTokenMiddleware, ticketController.UpdateTicket)
	r.DELETE("/ticket/deleteTicket/:ticketId", userModel.VerifyTokenMiddleware, ticketController.DeleteTicket)
}
