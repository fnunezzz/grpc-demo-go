package main

import (
	"log"
	"net"

	"github.com/fnunezzz/grpc-demo-go/internal/controller"
	"github.com/fnunezzz/grpc-demo-go/internal/repository"
	"github.com/fnunezzz/grpc-demo-go/internal/service"
	"github.com/fnunezzz/grpc-demo-go/internal/usecase"
	"google.golang.org/grpc"

	// "github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)


func main() {
	godotenv.Load()
    listen, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalln(err.Error())
	}
	server := grpc.NewServer()

	// api := app.Group("/api")
	dao := repository.InitConn()


	// Services
	appService := service.NewAppService(dao)
	
	// Usecases
	loginUserUseCase := usecase.NewLoginUserUseCase(dao)
	createUserUseCase := usecase.NewCreateUserUseCase(dao)

	// Routes
	controller.NewUserController(server, loginUserUseCase, createUserUseCase)
	controller.NewAppController(server, appService)
	
	err = server.Serve(listen)
	if err != nil {
		log.Fatalln(err.Error())
	}
}