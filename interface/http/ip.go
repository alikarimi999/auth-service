package httpserver

import (
	"fmt"
	"net/http"

	interfaces "github.com/alikarimi999/auth_service/interface"
	dto "github.com/alikarimi999/auth_service/interface/dto/http"
)

func (h *HttpServer) AddIp(ctx interfaces.ServerContext) {
	req := &dto.UpdateIpRequest{}
	err := ctx.Bind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
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

	err = h.App.AddIp(hashId(id), req.Ip, at, ctx.Request().Context())
	if err != nil {
		handlerErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, fmt.Sprintf("%s added", req.Ip))

}

func (h *HttpServer) RemoveIp(ctx interfaces.ServerContext) {
	req := &dto.UpdateIpRequest{}
	err := ctx.Bind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
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

	err = h.App.RemoveIp(hashId(id), req.Ip, at, ctx.Request().Context())
	if err != nil {
		handlerErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, fmt.Sprintf("%s removed", req.Ip))

}
