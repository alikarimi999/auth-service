package httpserver

import (
	"net/http"

	interfaces "github.com/alikarimi999/auth_service/interface"
	"github.com/alikarimi999/errors"
)

func handlerErr(ctx interfaces.ServerContext, err error) {
	switch errors.ErrorCode(err) {
	case errors.ErrNotFound:
		ctx.JSON(http.StatusNotFound, errors.ErrorMsg(err))
	case errors.ErrBadRequest:
		ctx.JSON(http.StatusBadRequest, errors.ErrorMsg(err))
	case errors.ErrInternal:
		ctx.JSON(http.StatusInternalServerError, errors.ErrorMsg(err))
	case errors.ErrForbidden:
		ctx.JSON(http.StatusForbidden, errors.ErrorMsg(err))
	default:
		ctx.JSON(http.StatusInternalServerError, errors.ErrorMsg(err))
	}
}
