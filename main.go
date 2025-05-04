package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/controller"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/controller/routes"
	"github.com/kevynlohan05/meu-primeiro-crud-go/src/model/service"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Init depedencies
	userService := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(userService)

	ticketService := service.NewTicketDomainService()
	ticketController := controller.NewTicketControllerInterface(ticketService)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController, ticketController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
