package app

import "github.com/billsbook/auth_service/domain"

type Application struct {
	repo domain.AuthRepo
}

func NewApp(repo domain.AuthRepo) *Application {
	return &Application{
		repo: repo,
	}
}
