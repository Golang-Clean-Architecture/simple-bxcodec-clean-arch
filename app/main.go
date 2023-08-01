package main

import (
	httpTodo "bxcodec-clean-arch/article/delivery/http"
	"bxcodec-clean-arch/article/usecase"
	"bxcodec-clean-arch/domain"
	repository "bxcodec-clean-arch/todo/repository/postgresql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	server         *gin.Engine
	todoService    usecase.TodoServiceImpl
	todoController httpTodo.TodoController
)

func init() {
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	// Init DB
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", viper.GetString("database.user"), viper.GetString("database.pass"), viper.GetString("database.host"), viper.GetString("database.port"), viper.GetString("database.name"))
	dbConn, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Migrate the schema
	dbConn.AutoMigrate(&domain.Todo{})

	// Check db connection

	fmt.Println("Server is running on port " + viper.GetString("server.address"))

	// Init Router

	server = gin.Default()
	basePath := server.Group("/v1")

	// Init Middleware

	// Init Repository
	todoRepo := repository.NewPostgresqlTodoRepo(dbConn)
	// Init Usecase
	todoService = *usecase.NewTodoService(todoRepo)
	// Init Handler
	todoController = httpTodo.NewTodoController(todoService)

	// Run Router
	todoController.RegisterTodoRoutes(basePath)
	log.Fatal(server.Run(":8080"))

}
