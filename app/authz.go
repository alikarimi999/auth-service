package app

import (
	"context"
	"fmt"

	"github.com/alikarimi999/errors"
	"github.com/billsbook/auth_service/domain"
)

func (a *Application) NewActore(actore domain.Actor, ctx context.Context) error {
	if len(actore.IPs()) >= int(a.maxIps) {
		return errors.Wrap(errors.ErrForbidden, errors.NewMesssage(fmt.Sprintf("only %d ips allowed", a.maxIps)))
	}

	c, error := a.repo.CountTokens(actore.UserId(), ctx)
	if error != nil {
		return error
	}
	if c >= a.maxTokens {
		return errors.Wrap(errors.ErrForbidden, errors.NewMesssage(fmt.Sprintf("only %d tokens allowed", a.maxTokens)))
	}

	return a.repo.SaveActore(actore, ctx)
}

// func (a *Application) AddPermission(id domain.ActoreID, perm *domain.Permission, ctx context.Context) error {
// 	return a.repo.AddPerm(id, perm, ctx)
// }

func (a *Application) HasAccess(id, ip string, t domain.ActoreType, res domain.Resource,
	act domain.Action, ctx context.Context) (bool, string, error) {

	actore, err := a.repo.GetActore(id, t, ctx)
	if err != nil {
		return false, "", errors.Wrap(errors.ErrNotFound, err)
	}

	for _, i := range actore.IPs() {
		if ip == i {
			perm, ok := actore.Permissions()[res]
			if !ok {
				return false, "", nil
			}

			switch act {
			case domain.None:
				return true, "", nil
			case domain.Read:
				return perm.Action == domain.Read || perm.Action == domain.Write, "", nil
			case domain.Write:
				return perm.Action == domain.Write, "", nil
			default:
				return false, "", errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("invalid action"))
			}
		}
	}
	return false, fmt.Sprintf("invalid ip"), nil
}

func (a *Application) GetUserActores(userId int64, ctx context.Context) ([]domain.Actor, error) {
	return a.repo.GetUserActores(userId, ctx)
}

func (a *Application) GetActore(id string, t domain.ActoreType, ctx context.Context) (domain.Actor, error) {
	return a.repo.GetActore(id, t, ctx)
}

func (a *Application) RemoveActore(id string, t domain.ActoreType, ctx context.Context) error {
	return a.repo.RemoveActore(id, t, ctx)
}
