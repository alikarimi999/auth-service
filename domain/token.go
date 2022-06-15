package domain

type APIToken struct {
	Id    ActoreID
	Perms map[Resource]*Permission
}

func (t *APIToken) ID() string {
	return "api_key_" + t.Id.String()
}

func (p *APIToken) Permissions() map[Resource]*Permission {
	return p.Perms
}

func (t *APIToken) AddPermission(p *Permission) {
	t.Perms[p.Resource] = p
}

func (t *APIToken) ChangePermission(res Resource, act Action) {
	t.Perms[res].Action = act
}

func (t *APIToken) Type() ActoreType {
	return Token
}

type TokenID struct {
	ApiKey string
}

func (i *TokenID) String() string {
	return i.ApiKey
}

func (*TokenID) IDType() ActoreIDType {
	return ActoreAPIToken
}
