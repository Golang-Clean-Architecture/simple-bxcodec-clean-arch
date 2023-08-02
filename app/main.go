package main

import (
	"bxcodec-clean-arch/domain"
	httpTodo "bxcodec-clean-arch/todo/delivery/http"
	repository "bxcodec-clean-arch/todo/repository/postgresql"
	"bxcodec-clean-arch/todo/usecase"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	server *gin.Engine
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
	todoService := usecase.NewTodoService(todoRepo)
	// Init Handler
	todoController := httpTodo.NewTodoController(todoService)

	// Run Router
	todoController.RegisterTodoRoutes(basePath)
	log.Fatal(server.Run(viper.GetString("server.address")))

}
