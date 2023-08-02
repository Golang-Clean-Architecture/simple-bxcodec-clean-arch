package domain

type Todo struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `json:"name" bson:"name"`
	Status string `json:"status" bson:"status"`
}

// here we create repository Contract
type TodoRepo interface {
	CreateTodo(*Todo) error
	GetTodo(*string) (*Todo, error)
	GetAll() ([]*Todo, error)
	UpdateTodo(*Todo) error
	DeleteTodo(*string) error
}

type TodoUsecase interface {
	CreateTodo(*Todo) error
	GetTodo(*string) (*Todo, error)
	GetAll() ([]*Todo, error)
	UpdateTodo(*Todo) error
	DeleteTodo(*string) error
}
