package main 

import (
    "github.com/depressed-kid/todo-app"
    "github.com/depressed-kid/todo-app/pkg/handler"
    "github.com/depressed-kid/todo-app/pkg/repository"
    "github.com/depressed-kid/todo-app/pkg/service"
    "log"
)

func main() {
    repos := repository.NewRepository()
    services := service.NewService(repos)
    handlers := handler.NewHandler(services)

    server := new(todo.HTTPServer)
    if err := server.Run("8000", handlers.InitRoutes()); err != nil {
	log.Fatalf("Error occured while running http server: %s", err.Error())
    } 
}



