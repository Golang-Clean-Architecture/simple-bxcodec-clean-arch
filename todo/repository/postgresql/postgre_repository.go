package repository

import (
	"bxcodec-clean-arch/domain"
	"errors"

	"gorm.io/gorm"
)

type PostgresqlTodoRepo struct {
	DB *gorm.DB
}

// New
// make a function that act like a constructor
func NewPostgresqlTodoRepo(dB *gorm.DB) domain.TodoRepo {
	return &PostgresqlTodoRepo{DB: dB}
}

// receiver function or more like classes/struct method in python/java
func (u *PostgresqlTodoRepo) CreateTodo(todo *domain.Todo) error {
	if todo.Name == "" {
		return errors.New("please fill todo name")
	}
	err := u.DB.Create(todo).Error
	return err
}

func (u *PostgresqlTodoRepo) GetTodo(name *string) (*domain.Todo, error) {
	var todo *domain.Todo
	err := u.DB.Where("name = ?", name).First(&todo).Error
	if err != nil {
		return nil, err
	}
	return todo, err
}

func (u *PostgresqlTodoRepo) GetAll() ([]*domain.Todo, error) {
	var todos []*domain.Todo
	err := u.DB.Find(&todos).Error
	if err != nil {
		return nil, err
	}
	if len(todos) == 0 {
		return nil, errors.New("document is empty")
	}
	return todos, nil
}

// Update All with same name
func (u *PostgresqlTodoRepo) UpdateTodo(todo *domain.Todo) error {
	err := u.DB.Model(&todo).Where("name = ?", todo.Name).Update("status", todo.Status).Error

	if err != nil {
		return errors.New("no matched document found for update")
	}
	return nil
}

func (u *PostgresqlTodoRepo) DeleteTodo(name *string) error {
	var todo *domain.Todo
	errFind := u.DB.Where("name = ?", name).First(&todo).Error
	if errFind != nil {
		return errors.New("no matched document found for delete")
	}
	u.DB.Delete(&domain.Todo{}, todo.ID)
	return nil
}
