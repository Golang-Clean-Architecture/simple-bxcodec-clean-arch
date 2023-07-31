package mocks

import (
	"bxcodec-clean-arch/domain"
	"errors"

	"github.com/stretchr/testify/mock"
)

type TodoRepositoryMock struct {
	Mock mock.Mock
}

func (repository *TodoRepositoryMock) GetTodo(name *string) (*domain.Todo, error) {
	arguments := repository.Mock.Called(name)
	if arguments.Get(0) == nil {
		return nil, errors.New("not found")
	} else {
		todo := arguments.Get(0).(domain.Todo)
		return &todo, nil
	}
}

func (repository *TodoRepositoryMock) GetAll() ([]*domain.Todo, error) {
	arguments := repository.Mock.Called()
	if arguments.Get(0) == nil {
		return nil, errors.New("document is empty")
	} else {
		todo := arguments.Get(0).([]*domain.Todo)
		return todo, nil
	}
}

func (repository *TodoRepositoryMock) CreateTodo(entitys *domain.Todo) error {
	arguments := repository.Mock.Called(entitys)
	if arguments.Get(0).(domain.Todo).Name == "" {
		return errors.New("please fill todo name")
	} else {
		return nil
	}
}

func (repository *TodoRepositoryMock) DeleteTodo(entitys *string) error {
	arguments := repository.Mock.Called(entitys)
	if arguments.Get(0).(domain.Todo).Name != *entitys {
		return errors.New("no matched document found for delete")
	} else {
		return nil
	}
}

func (repository *TodoRepositoryMock) UpdateTodo(entitys *domain.Todo) error {
	arguments := repository.Mock.Called(entitys)
	if arguments.Get(0).(domain.Todo).Name != entitys.Name {
		return errors.New("no matched document found for update")
	} else {
		return nil
	}
}
