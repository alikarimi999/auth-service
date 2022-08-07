package app

import "github.com/billsbook/auth_service/domain"

type Application struct {
	repo      domain.AuthRepo
	maxTokens int64
	maxIps    int64
}

func NewApp(repo domain.AuthRepo) *Application {
	return &Application{
		repo:      repo,
		maxTokens: 20,
		maxIps:    20,
	}
}
