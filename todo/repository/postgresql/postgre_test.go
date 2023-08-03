package repository_test

import (
	"bxcodec-clean-arch/domain"
	repository "bxcodec-clean-arch/todo/repository/postgresql"
	"database/sql"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestShouldGetTodoByName(t *testing.T) {
	db, mock, err := sqlmock.New()

	dialector := postgres.New(postgres.Config{
		Conn:       db,
		DriverName: "postgres",
	})
	dbGorm, _ := gorm.Open(dialector, &gorm.Config{})
	todoRepo := repository.NewPostgresqlTodoRepo(dbGorm)
	if err != nil {
		t.Fatal("an error " + err.Error() + " was not expected when opening a stub database connection")
	}
	defer db.Close()

	todoName := "Task 1" // in querying, this string turn into $1 = first arguments -> of the function

	query := `SELECT * FROM "todos" WHERE name = $1 ORDER BY "todos"."id" LIMIT 1`
	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(
			sqlmock.NewRows([]string{"ID", "name", "status"}).
				AddRow("5", "Task 1", "DONE"))

	todo, err := todoRepo.GetTodo(&todoName)

	assert.NoError(t, err)
	assert.NotNil(t, todo)

}

func TestShouldGetTodoByNameError(t *testing.T) {
	db, mock, err := sqlmock.New()

	dialector := postgres.New(postgres.Config{
		Conn:       db,
		DriverName: "postgres",
	})
	dbGorm, _ := gorm.Open(dialector, &gorm.Config{})
	todoRepo := repository.NewPostgresqlTodoRepo(dbGorm)
	if err != nil {
		t.Fatal("an error " + err.Error() + " was not expected when opening a stub database connection")
	}
	defer db.Close()

	todoName := "Task 1" // in querying, this string turn into $1 = first arguments -> of the function

	query := `SELECT * FROM "todos" WHERE name = $1 ORDER BY "todos"."id" LIMIT 1`
	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnError(sql.ErrConnDone)

	todo, err := todoRepo.GetTodo(&todoName)
	fmt.Println(todo)
	assert.Error(t, err)
	assert.Nil(t, todo)

}

func TestShouldCreateTodo(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	dialector := postgres.New(postgres.Config{
		Conn:       db,
		DriverName: "postgres",
	})
	dbGorm, _ := gorm.Open(dialector, &gorm.Config{})
	todoRepo := repository.NewPostgresqlTodoRepo(dbGorm)
	if err != nil {
		t.Fatal("an error " + err.Error() + " was not expected when opening a stub database connection")
	}
	defer db.Close()

	todoName := domain.Todo{
		ID:     1,
		Name:   "No Name",
		Status: "DONE",
	} // in querying, this string turn into $1 = first arguments -> of the function

	query := `INSERT INTO "todos" ("name","status", "id") VALUES ($1,$2,$3) RETURNING "id"`
	mock.ExpectBegin()
	mock.ExpectQuery(query).
		WithArgs("No Name", "DONE", 1).
		WillReturnError(nil)
	mock.ExpectCommit()
	err = todoRepo.CreateTodo(&todoName)

	assert.Nil(t, err)
}
