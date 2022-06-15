package app

import (
	"context"
	"errors"

	"github.com/billsbook/auth/domain"
)

func (a *Application) NewActore(actore domain.Actor, ctx context.Context) error {
	return a.repo.SaveActore(actore, ctx)
}

// func (a *Application) AddPermission(id domain.ActoreID, perm *domain.Permission, ctx context.Context) error {
// 	return a.repo.AddPerm(id, perm, ctx)
// }

func (a *Application) HasAccess(id domain.ActoreID, res domain.Resource, act domain.Action, ctx context.Context) (bool, error) {
	actore, err := a.repo.GetActore(id, ctx)
	if err != nil {
		return false, err
	}
	perm, ok := actore.Permissions()[res]
	if !ok {
		return false, nil
	}

	switch act {
	case domain.None:
		return true, nil
	case domain.Read:
		return perm.Action == domain.Read || perm.Action == domain.Write, nil
	case domain.Write:
		return perm.Action == domain.Write, nil
	default:
		return false, errors.New("")
	}

}

func (a *Application) Actore(id domain.ActoreID, ctx context.Context) (domain.Actor, error) {
	return a.repo.GetActore(id, ctx)
}
