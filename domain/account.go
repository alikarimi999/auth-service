package domain

type Account struct {
	Id    ActoreID
	User  string // User and UID is same
	Perms map[Resource]*Permission
}

func (a *Account) ID() string {
	return a.Id.String()
}

func (a *Account) Permissions() map[Resource]*Permission {
	return a.Perms
}

func (a *Account) AddPermission(p *Permission) {
	a.Perms[p.Resource] = p
}

func (a *Account) ChangePermission(res Resource, act Action) {
	a.Perms[res].Action = act
}

func (a *Account) Type() ActoreType {
	return Acc
}

type AccountID struct {
	User string
}

func (i *AccountID) String() string {
	return i.User
}

func (*AccountID) IDType() ActoreIDType {
	return ActoreAccount
}
