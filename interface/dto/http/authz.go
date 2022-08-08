package dto

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/alikarimi999/auth_service/domain"
	"github.com/alikarimi999/errors"
)

type Permission struct {
	Action string `json:"action"`
}

type CheckAccessRequest struct {
	Id       string `json:"id"`
	Ip       string `json:"ip"`
	Resource string `json:"resource"`
	Action   string `json:"action"`
}

func (r *CheckAccessRequest) Validate() error {
	if r.Id == "" {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("id is required"))
	}

	if r.Ip == "" {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("ip is required"))
	}

	if !isValidIP(r.Ip) {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("invalid ip"))
	}

	if r.Resource == "" {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("resource is required"))
	}

	if r.Action == "" {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("action is required"))
	}

	switch r.Resource {
	case "orders":
		if r.Action != "read" && r.Action != "write" {
			return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("action is required"))
		}

	default:
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("unknown resource"))
	}

	return nil
}

func ParseActoreId(id string) (string, domain.ActoreType, error) {
	if len(id) < 9 {
		return "", 0, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("invalid id"))
	}
	switch id[:8] {
	case "api_key_":
		return id[8:], domain.Token, nil
	default:
		return "", -1, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("unknown actor type"))
	}
}

type CheckAccessResp struct {
	Id        string `json:"id"`
	UserId    int64  `json:"user_id"`
	HasAccess bool   `json:"has_access"`
	Msg       string `json:"msg"`
}

type CreateActoreReq struct {
	UserId int64                  `json:"user_id"`
	Perms  map[string]*Permission `json:"permissions"`
	Ips    []string               `json:"ips"`
}

func (r *CreateActoreReq) Validate() error {
	if r.UserId == 0 {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("user_id is required"))
	}

	if len(r.Ips) == 0 {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("ips is required"))
	}

	if len(r.Perms) == 0 {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("at least one permission is required"))
	}

	// check if ips are valid
	for _, ip := range r.Ips {
		if !isValidIP(ip) {
			return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(fmt.Sprintf("invalid ip: %s", ip)))
		}
	}

	for r, p := range r.Perms {
		switch r {
		case "orders":
			if p.Action != "read" && p.Action != "write" {
				return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("action must be read or write"))
			}
		default:
			return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("unknown resource"))
		}
	}

	return nil
}

func (req *CreateActoreReq) Map(hashId string) (domain.Actor, error) {

	t := &domain.APIToken{
		Id:     hashId,
		Userid: req.UserId,
		Perms:  make(map[domain.Resource]*domain.Permission),
		Ips:    req.Ips,
	}

	for r, p := range req.Perms {
		res := Resource(r)
		if res == -1 {
			return nil, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("unknown resource"))
		}
		t.Perms[res] = &domain.Permission{
			Resource: res,
			Action:   Action(p.Action),
		}
	}

	return t, nil

}

type GetActoreResp struct {
	ID     string                 `json:"id"`
	UserId int64                  `json:"user_id"`
	Perms  map[string]*Permission `json:"permissions"`
	Ips    []string               `json:"ips"`
}

func GetActoreRespone(a domain.Actor) *GetActoreResp {

	resp := &GetActoreResp{
		ID:     a.ID()[9:],
		UserId: a.UserId(),
		Perms:  make(map[string]*Permission),
		Ips:    a.IPs(),
	}

	for r, p := range a.Permissions() {

		resp.Perms[r.String()] = &Permission{
			Action: p.Action.String(),
		}
	}
	return resp
}

func isValidIP(ip string) bool {
	if ip == "" {
		return false
	}
	ss := strings.Split(ip, ".")
	if len(ss) != 4 {
		return false
	}
	for _, s := range ss {
		i, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		if i < 0 || i > 255 {
			return false
		}
	}
	return true
}

type UpdateIpRequest struct {
	Id string `json:"id"`
	Ip string `json:"ip"`
}

func (r *UpdateIpRequest) Validate() error {
	if r.Id == "" {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("id is required"))
	}

	if r.Ip == "" {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("ip is required"))
	}

	if !isValidIP(r.Ip) {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("invalid ip"))
	}

	return nil
}
