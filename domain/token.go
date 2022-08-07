package domain

import "github.com/alikarimi999/errors"

type APIToken struct {
	Id     string
	Userid int64
	Perms  map[Resource]*Permission
	Ips    []string
}

func (t *APIToken) ID() string {
	return "api_key_" + t.Id
}

func (t *APIToken) UserId() int64 {
	return t.Userid
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

func (t *APIToken) AddIP(ip string) {
	t.Ips = append(t.Ips, ip)
}

func (t *APIToken) RemoveIp(ip string) error {
	for i, v := range t.Ips {
		if v == ip {
			t.Ips = append(t.Ips[:i], t.Ips[i+1:]...)
			return nil
		}
	}
	return errors.Wrap(errors.ErrNotFound, errors.NewMesssage("ip not found"))
}

func (t *APIToken) IPs() []string {
	return t.Ips
}
