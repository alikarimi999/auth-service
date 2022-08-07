package app

import (
	"context"
	"fmt"

	"github.com/alikarimi999/errors"
	"github.com/billsbook/auth_service/domain"
)

func (a *Application) AddIp(id, ip string, t domain.ActoreType, ctx context.Context) error {
	actore, err := a.repo.GetActore(id, t, ctx)
	if err != nil {
		return errors.Wrap(errors.ErrNotFound, errors.NewMesssage("actore not found"))
	}
	if len(actore.IPs()) >= int(a.maxIps) {
		return errors.Wrap(errors.ErrForbidden, errors.NewMesssage(fmt.Sprintf("only %d ips allowed", a.maxIps)))
	}

	for _, i := range actore.IPs() {
		if i == ip {
			return errors.Wrap(errors.ErrForbidden, errors.NewMesssage("ip already exists"))
		}
	}

	actore.AddIP(ip)
	if err := a.repo.UpdateActore(actore, ctx); err != nil {
		fmt.Println(err)
		return errors.Wrap(errors.ErrInternal, errors.NewMesssage("failed to update actore"))
	}
	return nil
}

func (a *Application) RemoveIp(id, ip string, t domain.ActoreType, ctx context.Context) error {
	actore, err := a.repo.GetActore(id, t, ctx)
	if err != nil {
		return errors.Wrap(errors.ErrNotFound, errors.NewMesssage("actore not found"))
	}

	if err := actore.RemoveIp(ip); err != nil {
		return err
	}
	if err := a.repo.UpdateActore(actore, ctx); err != nil {
		fmt.Println(err)
		return errors.Wrap(errors.ErrInternal, errors.NewMesssage("failed to update actore"))
	}
	return nil
}
