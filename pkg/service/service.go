package service

import "github.com/IKostarev/go-todo/pkg/repository"

type Authorization interface {

}

type TodoList interface{

}

type TodoItem interface {

}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) Service {
	return &Service{}  // ТУТ ПОЧЕМУ-ТО НЕ РАБОТАЕТ УКАЗАТЕЛЬ, ТАКЖЕ КАК И В РЕПОЗИТОРИИ, НУЖНО РАЗОБРАТЬСЯ
}