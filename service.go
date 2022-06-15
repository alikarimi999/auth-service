package main

import (
	"github.com/billsbook/auth_service/app"
	"github.com/billsbook/auth_service/infrastructure/database"
	"github.com/billsbook/auth_service/infrastructure/http"
	httpserver "github.com/billsbook/auth_service/interfaces/http"
)

func main() {
	repo := database.NewAuthzDB()

	app := app.NewApp(repo)
	si := httpserver.NewServer(app)
	router := http.NewRouter(si)
	router.Run(":8081")
}
