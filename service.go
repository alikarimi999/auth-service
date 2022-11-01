package main

import (
	"fmt"
	"os"

	"github.com/alikarimi999/auth_service/app"
	"github.com/alikarimi999/auth_service/infrastructure/database"
	"github.com/alikarimi999/auth_service/infrastructure/http"
	httpserver "github.com/alikarimi999/auth_service/interface/http"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// test()
	production()
}

func test() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root", "123", "localhost:3306", "auth_service")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to mysql")
	}

	repo := database.NewAuthzDB(db)

	app := app.NewApp(repo)
	si := httpserver.NewServer(app)
	router := http.NewRouter(si)
	fmt.Println("starting server on port 9091")
	router.Run(":9091")
}

func production() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), "auth_service")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to mysql")
	}

	repo := database.NewAuthzDB(db)

	app := app.NewApp(repo)
	si := httpserver.NewServer(app)
	router := http.NewRouter(si)
	fmt.Println("starting server on port 9091")
	router.Run(":8000")
}
