package routes

import (
	"github.com/gin-gonic/gin"
	controllerProject "github.com/kevynlohan05/meu-primeiro-crud-go/src/controller/project"
	controllerTicket "github.com/kevynlohan05/meu-primeiro-crud-go/src/controller/ticket"
	controllerUser "github.com/kevynlohan05/meu-primeiro-crud-go/src/controller/user"
	userModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user"
)

// InitRoutes registers all application routes with proper middleware and handlers
func InitRoutes(
	r *gin.RouterGroup,
	userController controllerUser.UserControllerInterface,
	ticketController controllerTicket.TicketControllerInterface,
	projectController controllerProject.ProjectControllerInterface,
) {
	// --- User routes ---
	r.POST("/user/login", userController.LoginUser)

	r.POST("/user/createUser", userController.CreateUser)
	r.GET("/user/getUserById/:userId", userModel.VerifyTokenMiddleware, userModel.AdminOnlyMiddleware, userController.FindUserById)
	r.GET("/user/getUserByEmail/:userEmail", userModel.VerifyTokenMiddleware, userModel.AdminOnlyMiddleware, userController.FindUserByEmail)
	r.PUT("/user/updateUser/:userId", userModel.VerifyTokenMiddleware, userModel.AdminOnlyMiddleware, userController.UpdateUser)
	r.DELETE("/user/deleteUser/:userId", userModel.VerifyTokenMiddleware, userModel.AdminOnlyMiddleware, userController.DeleteUser)

	// --- Ticket routes ---
	r.POST("/ticket/createTicket", userModel.VerifyTokenMiddleware, ticketController.CreateTicket)
	r.GET("/ticket/getTicketById/:ticketId", userModel.VerifyTokenMiddleware, userModel.AdminOnlyMiddleware, ticketController.FindTicketById)
	r.GET("/ticket/getAllTicketsByEmail/:ticketEmail", userModel.VerifyTokenMiddleware, ticketController.FindAllTicketsByEmail)
	r.GET("/ticket/getAllTicketsByEmailAndStatus/:ticketEmail/:ticketStatus", userModel.VerifyTokenMiddleware, ticketController.FindAllTicketsByEmailAndStatus)
	r.GET("/ticket/getAllTickets", userModel.VerifyTokenMiddleware, userModel.AdminOnlyMiddleware, ticketController.FindAllTickets)
	r.PUT("/ticket/updateTicket/:ticketId", userModel.VerifyTokenMiddleware, userModel.AdminOnlyMiddleware, ticketController.UpdateTicket)
	r.PUT("/ticket/updateComment/:ticketId/:commentId", userModel.VerifyTokenMiddleware, ticketController.UpdateComment)
	r.POST("/ticket/addComment/:ticketId", userModel.VerifyTokenMiddleware, ticketController.AddComment)
	r.DELETE("/ticket/deleteComment/:ticketId/:commentId", userModel.VerifyTokenMiddleware, ticketController.DeleteComment)
	r.DELETE("/ticket/deleteTicket/:ticketId", userModel.VerifyTokenMiddleware, userModel.AdminOnlyMiddleware, ticketController.DeleteTicket)

	// --- Project routes ---
	r.POST("/project/createProject", userModel.VerifyTokenMiddleware, userModel.AdminOnlyMiddleware, projectController.CreateProject)
	r.GET("/project/getProjectById/:projectId", userModel.VerifyTokenMiddleware, projectController.FindProjectById)
	r.GET("/project/getProjectByName/:projectName", userModel.VerifyTokenMiddleware, projectController.FindProjectByName)
	r.GET("/project/getAllProjects", userModel.VerifyTokenMiddleware, projectController.FindAllProjects)
	r.PUT("/project/updateProject/:projectId", userModel.VerifyTokenMiddleware, userModel.AdminOnlyMiddleware, projectController.UpdateProject)
	r.DELETE("/project/deleteProject/:projectId", userModel.VerifyTokenMiddleware, userModel.AdminOnlyMiddleware, projectController.DeleteProject)
}
