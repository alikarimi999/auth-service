package dto

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/billsbook/auth_service/domain"
)

type ApiToken struct {
	Id          string
	UserId      int64
	Permissions *perms
	Ips         *IPs
}

func NewAPIToken(dt *domain.APIToken) *ApiToken {
	var ips IPs
	if dt.Ips != nil {
		ips = IPs(dt.Ips)
	} else {
		ips = IPs{}
	}

	t := &ApiToken{
		Id:     dt.Id,
		UserId: dt.UserId(),
		Permissions: &perms{
			PS: make(map[int]*Permission),
		},
		Ips: &ips,
	}

	if dt.Perms != nil {
		for r, p := range dt.Perms {
			t.Permissions.PS[int(r)] = &Permission{
				Resource: p.Resource.String(),
				Action:   p.Action.String(),
			}
		}
	}
	return t
}

func (a *ApiToken) Map() *domain.APIToken {
	var ips IPs
	if a.Ips != nil {
		ips = *a.Ips
	} else {
		ips = IPs{}
	}
	t := &domain.APIToken{
		Id:     a.Id,
		Userid: a.UserId,
		Perms:  make(map[domain.Resource]*domain.Permission),
		Ips:    []string(ips),
	}

	if a.Permissions != nil {
		for r, p := range a.Permissions.PS {
			t.Perms[domain.Resource(r)] = &domain.Permission{
				Resource: domain.ParseResource(p.Resource),
				Action:   domain.ParseAction(p.Action),
			}
		}
	}
	return t
}

type perms struct {
	PS map[int]*Permission `json:"permissions"`
}

type Permission struct {
	Resource string `json:"resource"`
	Action   string `json:"action"`
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

type IPs []string

func (i *IPs) Value() (driver.Value, error) {
	bytes, err := json.Marshal(i)
	return string(bytes), err
}

func (i *IPs) Scan(value interface{}) error {
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
	return json.Unmarshal(bytes, i)

}
