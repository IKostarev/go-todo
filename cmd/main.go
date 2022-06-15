package main

import (
	"log"

	"github.com/IKostarev/go-todo"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("Error start server is: %s", err.Error())
	}	
}