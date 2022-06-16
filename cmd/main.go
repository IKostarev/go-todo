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

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5436",
		Username: "postgres",
		Password: "qwerty",
		DBName:   "postgres",
		SSLMode:  "disable",
	})

	if err != nil {
		log.Fatalf("Error to init db is: %s", err.Error())
	}

	repos := repository.NewRepository(db)
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