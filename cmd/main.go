package main

import (
	"log"

	"github.com/IKostarev/go-todo"
	"github.com/IKostarev/go-todo/pkg/handler"
	"github.com/IKostarev/go-todo/pkg/repository"
	"github.com/IKostarev/go-todo/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error start server is: %s", err.Error())
	}	
}