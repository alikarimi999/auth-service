package httpserver

import (
	"net/http"

	"github.com/alikarimi999/errors"
	interfaces "github.com/billsbook/auth_service/interface"
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
