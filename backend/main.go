package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/database/mysql"
	controllerProject "github.com/kevynlohan05/meu-primeiro-crud-go/src/controller/project"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/controller/routes"
	controllerTicket "github.com/kevynlohan05/meu-primeiro-crud-go/src/controller/ticket"
	controllerUser "github.com/kevynlohan05/meu-primeiro-crud-go/src/controller/user"
	repositoryProject "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/projects/repository"
	projectService "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/projects/service"
	repositoryTicket "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket/repository"
	ticketService "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket/service"
	repositoryUser "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user/repository"
	userService "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user/service"
)

func main() {
	// Carrega variáveis de ambiente
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Conexão MySQL
	database, err := mysql.NewMySQLConnection()
	if err != nil {
		log.Fatalf("Error connecting to MySQL, error=%s \n", err.Error())
		return
	}

	// Inicializa repositórios e serviços com *sql.DB
	repoUser := repositoryUser.NewUserRepository(database)
	userServiceInstance := userService.NewUserDomainService(repoUser)
	userController := controllerUser.NewUserControllerInterface(userServiceInstance)

	repoTicket := repositoryTicket.NewTicketRepository(database)
	ticketServiceInstance := ticketService.NewTicketDomainService(userServiceInstance, repoTicket)
	ticketController := controllerTicket.NewTicketControllerInterface(ticketServiceInstance)

	repoProject := repositoryProject.NewProjectRepository(database)
	projectServiceInstance := projectService.NewProjectService(repoProject)
	projectController := controllerProject.NewProjectControllerInterface(projectServiceInstance)

	// Setup Gin com CORS
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Inicializa rotas
	routes.InitRoutes(&router.RouterGroup, userController, ticketController, projectController)

	// Inicia servidor
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
