package domain

import (
	"context"
)

type Actor interface {
	ID() string
	Permissions() map[Resource]*Permission
	AddPermission(p *Permission)
	ChangePermission(res Resource, act Action)
	Type() ActoreType
}

type AuthRepo interface {
	GetActore(id ActoreID, ctx context.Context) (Actor, error)
	// AddPerm(id ActoreID, perm *Permission, ctx context.Context) error
	SaveActore(a Actor, ctx context.Context) error
}

type ActoreIDType int

const (
	ActoreAPIToken ActoreIDType = iota
	ActoreAccount
)

type ActoreID interface {
	String() string
	IDType() ActoreIDType
}
