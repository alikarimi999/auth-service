package domain

import (
	"context"
)

type Actor interface {
	ID() string
	UserId() int64
	Permissions() map[Resource]*Permission
	IPs() []string
	AddPermission(p *Permission)
	AddIP(ip string)
	RemoveIp(ip string) error
	ChangePermission(res Resource, act Action)
	Type() ActoreType
}

type AuthRepo interface {
	GetUserActores(userId int64, ctx context.Context) ([]Actor, error)
	GetActore(id string, t ActoreType, ctx context.Context) (Actor, error)
	RemoveActore(id string, t ActoreType, ctx context.Context) error
	// AddPerm(id ActoreID, perm *Permission, ctx context.Context) error
	SaveActore(a Actor, ctx context.Context) error
	UpdateActore(a Actor, ctx context.Context) error
	CountTokens(userId int64, ctx context.Context) (int64, error)
}
