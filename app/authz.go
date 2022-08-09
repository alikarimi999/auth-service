package app

import (
	"context"
	"fmt"

	"github.com/alikarimi999/auth_service/domain"
	"github.com/alikarimi999/errors"
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

func (a *Application) GetUserActores(userId int64, ctx context.Context) ([]domain.Actor, error) {
	return a.repo.GetUserActores(userId, ctx)
}

func (a *Application) GetActore(id string, t domain.ActoreType, ctx context.Context) (domain.Actor, error) {
	return a.repo.GetActore(id, t, ctx)
}

func (a *Application) RemoveActore(id string, t domain.ActoreType, ctx context.Context) error {
	return a.repo.RemoveActore(id, t, ctx)
}
