package app

import "github.com/billsbook/auth/domain"

type Application struct {
	repo domain.AuthRepo
}

func NewApp(repo domain.AuthRepo) *Application {
	return &Application{
		repo: repo,
	}
}
