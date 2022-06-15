package httpserver

import (
	"errors"
	"net/http"

	"github.com/billsbook/auth/interfaces"
	dto "github.com/billsbook/auth/interfaces/dto/http"
)

func (h *httpServer) CheckAccess(ctx interfaces.ServerContext) {
	req := &dto.CheckAccessRequest{}
	err := ctx.Bind(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "")
		return
	}

	ok, err := h.App.HasAccess(dto.ActoreID(req.ID), dto.Resource(req.Resource),
		dto.Action(req.Action), ctx.Request().Context())
	if err != nil {
		ctx.JSON(http.StatusNotFound, "")
		return
	}
	ctx.JSON(http.StatusAccepted, dto.CheckAccessResp{HasAccess: ok})
}

func (h *httpServer) AddActore(ctx interfaces.ServerContext) {
	req := &dto.CreateActoreReq{}
	err := ctx.Bind(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "")
		return
	}

	actore := req.Map()
	if actore == nil {
		ctx.JSON(http.StatusBadRequest, "")
		return
	}
	err = h.App.NewActore(req.Map(), ctx.Request().Context())

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errors.New(err.Error()))
		return
	}
	ctx.JSON(http.StatusCreated, "created successfully")
}

func (h *httpServer) GetActore(ctx interfaces.ServerContext) {

	id := ctx.Param("id")

	actore, err := h.App.Actore(dto.ActoreID(id), ctx.Request().Context())
	if err != nil {
		ctx.JSON(http.StatusNotFound, "")
		return
	}
	ctx.JSON(http.StatusOK, dto.GetActoreRespone(actore))

}
