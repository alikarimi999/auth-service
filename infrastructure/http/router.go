package http

import (
	"github.com/billsbook/auth_service/interfaces"
	middlewares "github.com/billsbook/common/middleware"
	"github.com/gin-gonic/gin"
)

type Router struct {
	gin *gin.Engine
}

func (r *Router) Run(addr ...string) {
	r.gin.Run(addr...)
}

func NewRouter(si interfaces.ServerInterface) *Router {
	r := gin.New()
	v1 := r.Group("/v1")
	authz := v1.Group("/authz", middlewares.TenantHandler())
	authz.POST("/check_access", func(ctx *gin.Context) { si.CheckAccess(NewContext(ctx)) })
	authz.POST("/new_actore", func(ctx *gin.Context) { si.AddActore(NewContext(ctx)) })
	authz.GET("/get_actore/:id", func(ctx *gin.Context) { si.GetActore(NewContext(ctx)) })
	return &Router{gin: r}

}
