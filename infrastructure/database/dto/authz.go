package dto

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/billsbook/auth/domain"
	"gorm.io/gorm"
)

type APIToken struct {
	gorm.Model
	Key         string
	Permissions *perms
}

func NewAPIToken(dt *domain.APIToken) *APIToken {
	t := &APIToken{
		Key: dt.Id.String(),
		Permissions: &perms{
			PS: make(map[int]*Permission),
		},
	}

	for r, p := range dt.Perms {
		t.Permissions.PS[int(r)] = &Permission{
			Resource: int(p.Resource),
			Action:   int(p.Action),
		}
	}
	return t
}

func (a *APIToken) Map() *domain.APIToken {
	t := &domain.APIToken{
		Id:    &domain.TokenID{ApiKey: a.Key},
		Perms: make(map[domain.Resource]*domain.Permission),
	}

	for r, p := range a.Permissions.PS {
		t.Perms[domain.Resource(r)] = &domain.Permission{
			Resource: domain.Resource(p.Resource),
			Action:   domain.Action(p.Action),
		}
	}
	return t
}

type perms struct {
	PS map[int]*Permission `json:"permissions"`
}

type Permission struct {
	Resource int `json:"resource"`
	Action   int `json:"action"`
}

func (p *perms) Value() (driver.Value, error) {
	bytes, err := json.Marshal(p)
	return string(bytes), err
}

func (p *perms) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	var bytes []byte
	switch t := value.(type) {
	case string:
		bytes = []byte(t)
	case []byte:
		bytes = t
	}
	return json.Unmarshal(bytes, p)

}

func (*Permission) GormDataType() string {
	return "json"
}
