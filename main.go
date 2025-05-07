package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/configuration/database/mongodb"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/controller"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/controller/routes"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model/repository"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model/service"
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
	repoUser := repository.NewUserRepository(database)
	userService := service.NewUserDomainService(repoUser)
	userController := controller.NewUserControllerInterface(userService)

	repoTicket := repository.NewTicketRepository(database)
	ticketService := service.NewTicketDomainService(repoTicket)
	ticketController := controller.NewTicketControllerInterface(ticketService)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController, ticketController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
