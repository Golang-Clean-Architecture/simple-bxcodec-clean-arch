package repository_test

import (
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
