package main 

import (
    "github.com/depressed-kid/todo-app"
    "github.com/depressed-kid/todo-app/pkg/handler"
    "log"
)

func main() {
    handlers := new(handler.Handler)

    server := new(todo.HTTPServer)
    if err := server.Run("8000", handlers.InitRoutes()); err != nil {
	log.Fatalf("Error occured while running http server: %s", err.Error())
    } 
}



