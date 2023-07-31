package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
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
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", viper.GetString("database.user"), viper.GetString("database.pass"), viper.GetString("database.host"), viper.GetString("database.port"), viper.GetString("database.name"))
	dbConn, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	routerApp := echo.New()
	routerApp.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	routerApp.Logger.Fatal(routerApp.Start(viper.GetString("server.address")))
	fmt.Println("Server is running on port " + viper.GetString("server.address"))
}
