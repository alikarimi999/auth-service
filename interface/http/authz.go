package httpserver

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alikarimi999/errors"
	"github.com/billsbook/auth_service/domain"
	interfaces "github.com/billsbook/auth_service/interface"
	dto "github.com/billsbook/auth_service/interface/dto/http"
)

func (h *HttpServer) CheckAccess(ctx interfaces.ServerContext) {
	req := &dto.CheckAccessRequest{}
	err := ctx.Bind(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "")
		return
	}

	if err := req.Validate(); err != nil {
		handlerErr(ctx, err)
		return
	}

	id, at, err := dto.ParseActoreId(req.Id)
	if err != nil {
		handlerErr(ctx, err)
		return
	}

	actore, err := h.App.GetActore(hashId(id), at, ctx.Request().Context())
	if err != nil {
		handlerErr(ctx, err)
		return
	}

	resp := dto.CheckAccessResp{
		Id:     req.Id,
		UserId: actore.UserId(),
	}

	for _, i := range actore.IPs() {
		if req.Ip == i {
			perm, ok := actore.Permissions()[domain.ParseResource(req.Resource)]
			if !ok {
				handlerErr(ctx, errors.Wrap(errors.ErrNotFound, errors.NewMesssage("resource not found")))
				return
			}

			switch domain.ParseAction(req.Action) {
			case domain.None:
				resp.HasAccess = true
				ctx.JSON(http.StatusOK, resp)
				return
			case domain.Read:
				resp.HasAccess = perm.Action == domain.Read || perm.Action == domain.Write
				ctx.JSON(http.StatusOK, resp)
				return
			case domain.Write:
				resp.HasAccess = perm.Action == domain.Write
				ctx.JSON(http.StatusOK, resp)
				return
			default:
				handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("invalid action")))
				return
			}
		}
	}

	resp.HasAccess = false
	resp.Msg = "ip not found"
	ctx.JSON(http.StatusOK, resp)
	return
}

func (h *HttpServer) NewToken(ctx interfaces.ServerContext) {
	req := &dto.CreateActoreReq{}
	err := ctx.Bind(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := req.Validate(); err != nil {
		handlerErr(ctx, err)
		return
	}

	id := generateId(32)
	hid := hashId(id)

	actore, err := req.Map(hid)
	if err != nil {
		handlerErr(ctx, err)
		return
	}
	resp := struct {
		Id      string `json:"id,omitempty"`
		UserId  int64  `json:"user_id,omitempty"`
		Created bool   `json:"created,omitempty"`
		Msg     string `json:"message,omitempty"`
	}{}
	err = h.App.NewActore(actore, ctx.Request().Context())
	if err != nil {
		handlerErr(ctx, err)
		return
	}
	resp.Created = true
	resp.Id = fmt.Sprintf("api_key_%s", id)
	resp.UserId = req.UserId
	ctx.JSON(http.StatusCreated, resp)
	return
}

func (h *HttpServer) GetActore(ctx interfaces.ServerContext) {
	id := ctx.Param("id")
	id, at, err := dto.ParseActoreId(id)
	if err != nil {
		handlerErr(ctx, err)
		return
	}
	actore, err := h.App.GetActore(hashId(id), at, ctx.Request().Context())
	if err != nil {
		ctx.JSON(http.StatusNotFound, "")
		return
	}
	ctx.JSON(http.StatusOK, dto.GetActoreRespone(actore))
	return

}

func (h *HttpServer) GetUserActores(ctx interfaces.ServerContext) {
	userId, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		handlerErr(ctx, err)
		return
	}
	actores, err := h.App.GetUserActores(int64(userId), ctx.Request().Context())
	if err != nil {
		ctx.JSON(http.StatusNotFound, "")
		return
	}

	rs := []*dto.GetActoreResp{}

	for _, actore := range actores {
		rs = append(rs, dto.GetActoreRespone(actore))
	}
	ctx.JSON(http.StatusOK, rs)
	return
}

func (h *HttpServer) RemoveToken(ctx interfaces.ServerContext) {
	id := ctx.Param("id")
	id, at, err := dto.ParseActoreId(id)
	if err != nil {
		handlerErr(ctx, err)
		return
	}
	err = h.App.RemoveActore(hashId(id), at, ctx.Request().Context())
	if err != nil {
		handlerErr(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, fmt.Sprintf("token %s removed", id))
	return
}
