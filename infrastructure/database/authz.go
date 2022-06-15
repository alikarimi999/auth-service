package database

import (
	"context"
	"errors"

	"github.com/billsbook/auth/domain"
	"github.com/billsbook/auth/infrastructure/database/dto"
	"github.com/billsbook/common"
)

type AuthZDB struct {
	store common.DBStore
}

func NewAuthzDB() domain.AuthRepo {
	return &AuthZDB{store: *common.NewStore()}
}

func (az *AuthZDB) GetActore(id domain.ActoreID, ctx context.Context) (domain.Actor, error) {
	db, err := az.store.GetDB(ctx, common.Auth)
	if err != nil {
		return nil, err
	}
	switch id.IDType() {
	case domain.ActoreAPIToken:
		token := dto.APIToken{}
		db = db.Where("`key` = ?", id.String()).First(&token)
		if db.Error != nil {
			return nil, db.Error
		}
		return token.Map(), db.Error
	default:
		return nil, errors.New("")
	}
}

// func (az *AuthZDB) AddPerm(id domain.ActoreID, perm *domain.Permission, ctx context.Context) error {

// 	a, err := az.GetActore(id, ctx)
// 	if err != nil {
// 		return err
// 	}
// 	a.AddPermission(perm)
// 	return az.SaveActore(a, ctx)
// }

func (az *AuthZDB) SaveActore(a domain.Actor, ctx context.Context) error {

	db, err := az.store.GetDB(ctx, common.Auth)
	if err != nil {
		return err
	}

	switch a.Type() {
	case domain.Token:
		t := dto.NewAPIToken(a.(*domain.APIToken))
		db = db.Create(t)
		return db.Error
	default:
		return errors.New("")
	}

}
