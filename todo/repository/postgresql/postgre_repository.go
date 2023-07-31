package repository

import (
	"bxcodec-clean-arch/domain"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoRepo struct {
	todoCollection *mongo.Collection
	ctx            context.Context
}

// New
// make a function that act like a constructor
func New(mongo *mongo.Collection, ctx context.Context) *TodoRepo {
	return &TodoRepo{mongo, ctx}
}

// receiver function or more like classes/struct method in python/java
func (u *TodoRepo) CreateTodo(todo *domain.Todo) error {
	if todo.Name == "" {
		return errors.New("please fill todo name")
	}
	_, err := u.todoCollection.InsertOne(u.ctx, todo)
	return err
}

func (u *TodoRepo) GetTodo(name *string) (*domain.Todo, error) {
	var todo *domain.Todo
	query := bson.D{bson.E{Key: "name", Value: name}}
	err := u.todoCollection.FindOne(u.ctx, query).Decode(&todo)
	return todo, err
}

func (u *TodoRepo) GetAll() ([]*domain.Todo, error) {
	var todos []*domain.Todo
	cursor, err := u.todoCollection.Find(u.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(u.ctx) {
		var todo *domain.Todo
		err := cursor.Decode(&todo)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(u.ctx)

	if len(todos) == 0 {
		return nil, errors.New("document is empty")
	}
	return todos, nil
}

func (u *TodoRepo) UpdateTodo(todo *domain.Todo) error {
	filter := bson.D{bson.E{Key: "name", Value: todo.Name}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "name", Value: todo.Name}, bson.E{Key: "status", Value: todo.Status}}}}
	result, _ := u.todoCollection.UpdateOne(u.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}

func (u *TodoRepo) DeleteTodo(name *string) error {
	filter := bson.D{bson.E{Key: "name", Value: name}}
	result, _ := u.todoCollection.DeleteOne(u.ctx, filter)
	if result.DeletedCount == 0 {
		return errors.New("no matched document found for delete")
	}
	return nil
}
