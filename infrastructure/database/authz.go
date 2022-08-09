package database

import (
	"context"

	"github.com/alikarimi999/auth_service/domain"
	"github.com/alikarimi999/auth_service/infrastructure/database/dto"
	"github.com/alikarimi999/errors"
	"gorm.io/gorm"
)

type AuthZDB struct {
	db *gorm.DB
}

func NewAuthzDB(db *gorm.DB) domain.AuthRepo {
	return &AuthZDB{db: db}
}

func (az *AuthZDB) GetActore(id string, t domain.ActoreType, ctx context.Context) (domain.Actor, error) {
	switch t {
	case domain.Token:
		token := dto.ApiToken{}
		if err := az.db.Where(" id = ?", id).First(&token).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, errors.Wrap(errors.ErrNotFound)
			}
			return nil, err
		}
		return token.Map(), nil
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

	switch a.Type() {
	case domain.Token:
		t := dto.NewAPIToken(a.(*domain.APIToken))
		if err := az.db.Create(t).Error; err != nil {
			return err
		}
		return nil
	default:
		return errors.New("")
	}

}

func (az *AuthZDB) UpdateActore(a domain.Actor, ctx context.Context) error {

	switch a.Type() {
	case domain.Token:
		t := dto.NewAPIToken(a.(*domain.APIToken))
		if err := az.db.Save(t).Error; err != nil {
			return err
		}
		return nil
	default:
		return errors.New("")
	}

}

func (az *AuthZDB) GetUserActores(userId int64, ctx context.Context) ([]domain.Actor, error) {
	var actores []*dto.ApiToken
	if err := az.db.Where("user_id = ?", userId).Find(&actores).Error; err != nil {
		return nil, err
	}

	var actors []domain.Actor
	for _, a := range actores {
		actors = append(actors, a.Map())
	}

	return actors, nil
}

func (az *AuthZDB) RemoveActore(id string, t domain.ActoreType, ctx context.Context) error {
	switch t {
	case domain.Token:
		return az.db.Delete(&dto.ApiToken{Id: id}).Error
	default:
		return errors.New("")
	}
}

// count the number of tokens for a user
func (az *AuthZDB) CountTokens(userId int64, ctx context.Context) (int64, error) {
	var count int64
	if err := az.db.Model(&dto.ApiToken{}).Where("user_id = ?", userId).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
