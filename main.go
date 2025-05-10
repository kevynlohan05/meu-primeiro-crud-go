package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/database/mongodb"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/controller/routes"
	controllerTicket "github.com/kevynlohan05/meu-primeiro-crud-go/src/controller/ticket"
	controllerUser "github.com/kevynlohan05/meu-primeiro-crud-go/src/controller/user"
	repositoryTicket "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket/repository"
	ticketService "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket/service"
	repositoryUser "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user/repository"
	userService "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/user/service"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := mongodb.NewMongoDbConnection(context.Background())
	if err != nil {
		log.Fatalf("Error connecting to MongoDB, error=%s \n", err.Error())
		return
	}

	//Init depedencies
	repoUser := repositoryUser.NewUserRepository(database)
	userServiceInstace := userService.NewUserDomainService(repoUser)
	userController := controllerUser.NewUserControllerInterface(userServiceInstace)

	repoTicket := repositoryTicket.NewTicketRepository(database)
	ticketServiceInstance := ticketService.NewTicketDomainService(repoTicket)
	ticketController := controllerTicket.NewTicketControllerInterface(ticketServiceInstance)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController, ticketController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
