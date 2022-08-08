package main

import (
	"fmt"

	"github.com/alikarimi999/auth_service/app"
	"github.com/alikarimi999/auth_service/infrastructure/database"
	"github.com/alikarimi999/auth_service/infrastructure/http"
	httpserver "github.com/alikarimi999/auth_service/interface/http"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root", "123", "localhost:3306", "auth_service")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect to %s", dsn))
	}

	repo := database.NewAuthzDB(db)

	app := app.NewApp(repo)
	si := httpserver.NewServer(app)
	router := http.NewRouter(si)
	router.Run(":9091")
}
