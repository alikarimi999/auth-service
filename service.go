package main

import (
	"github.com/billsbook/auth/app"
	"github.com/billsbook/auth/infrastructure/database"
	"github.com/billsbook/auth/infrastructure/http"
	httpserver "github.com/billsbook/auth/interfaces/http"
)

func main() {
	repo := database.NewAuthzDB()

	app := app.NewApp(repo)
	si := httpserver.NewServer(app)
	router := http.NewRouter(si)
	router.Run(":8081")
}
