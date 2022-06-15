package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExtContext struct {
	ctx *gin.Context
}

func (ec *ExtContext) Param(p string) string {
	return ec.ctx.Param(p)
}
func (ec *ExtContext) Bind(i interface{}) error {
	return ec.ctx.Bind(i)
}
func (ec *ExtContext) JSON(code int, obj interface{}) {
	ec.ctx.JSON(code, obj)
}
func (ec *ExtContext) Request() *http.Request {
	return ec.ctx.Request
}
func (ec *ExtContext) GetKey(key string) (value interface{}, exists bool) {
	return ec.ctx.Get(key)
}

func (ec *ExtContext) SetKey(key string, value interface{}) {
	ec.ctx.Set(key, value)
}

func NewContext(ctx *gin.Context) *ExtContext {
	return &ExtContext{ctx: ctx}
}
