package main

import (
	"log"

	"github.com/IKostarev/go-todo"
	"github.com/IKostarev/go-todo/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error start server is: %s", err.Error())
	}	
}