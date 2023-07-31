package usecase

import (
	"bxcodec-clean-arch/domain"
)

type TodoServiceImpl struct {
	repo domain.TodoRepo
}

// make a function that act like a constructor
func NewTodoService(todoRepo domain.TodoRepo) *TodoServiceImpl {
	return &TodoServiceImpl{
		repo: todoRepo,
	}
}

// receiver function or more like classes/struct method in python/java
func (u *TodoServiceImpl) CreateTodo(todo *domain.Todo) error {
	err := u.repo.CreateTodo(todo)
	return err
}

func (u *TodoServiceImpl) GetTodo(name *string) (*domain.Todo, error) {
	todo, err := u.repo.GetTodo(name)
	return todo, err
}

func (u *TodoServiceImpl) GetAll() ([]*domain.Todo, error) {
	todos, err := u.repo.GetAll()
	return todos, err
}

func (u *TodoServiceImpl) UpdateTodo(todo *domain.Todo) error {
	err := u.repo.UpdateTodo(todo)
	return err
}

func (u *TodoServiceImpl) DeleteTodo(name *string) error {
	err := u.repo.DeleteTodo(name)
	return err
}
