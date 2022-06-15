package repository

import "github.com/IKostarev/go-todo/pkg/service"


type Authorization interface {

}

type TodoList interface{

}

type TodoItem interface {

}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(services *service.Service) Repository {
	return Repository{}
}