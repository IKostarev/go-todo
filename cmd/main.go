package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

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

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error load env: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
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
