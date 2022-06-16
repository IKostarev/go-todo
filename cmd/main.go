package main

import (
	"log"

	"github.com/IKostarev/go-todo"
	"github.com/IKostarev/go-todo/pkg/handler"
	"github.com/IKostarev/go-todo/pkg/repository"
	"github.com/IKostarev/go-todo/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initCofig(); err != nil {
		log.Fatalf("Error read config file is: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services) 

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error start server is: %s", err.Error())
	}	
}

func initCofig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}