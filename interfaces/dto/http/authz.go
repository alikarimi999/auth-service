package dto

import (
	"github.com/billsbook/auth/domain"
)

type Permission struct {
	Action string `json:"action"`
}

type CheckAccessRequest struct {
	ID       string `json:"id"`
	Resource string `json:"resource"`
	Action   string `json:"action"`
}

type CheckAccessResp struct {
	HasAccess bool `json:"has_access"`
}

type CreateActoreReq struct {
	ID    string                 `json:"id"`
	Perms map[string]*Permission `json:"permissions"`
}

func (req *CreateActoreReq) Map() domain.Actor {
	switch req.ID[:8] {
	case "api_key_":
		t := &domain.APIToken{
			Id:    &domain.TokenID{ApiKey: req.ID[8:]},
			Perms: make(map[domain.Resource]*domain.Permission),
		}

		for r, p := range req.Perms {
			res := Resource(r)
			if res == -1 {
				return nil
			}
			t.Perms[res] = &domain.Permission{
				Resource: res,
				Action:   Action(p.Action),
			}
		}
		return t
	default:
		return nil
	}

}

type GetActoreResp struct {
	ID    string                 `json:"id"`
	Perms map[string]*Permission `json:"permissions"`
}

func GetActoreRespone(a domain.Actor) *GetActoreResp {

	resp := &GetActoreResp{
		ID:    a.ID(),
		Perms: make(map[string]*Permission),
	}

	for r, p := range a.Permissions() {

		resp.Perms[r.String()] = &Permission{
			Action: p.Action.String(),
		}
	}
	return resp
}
