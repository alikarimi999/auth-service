package http

import (
	httpserver "github.com/alikarimi999/auth_service/interface/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	gin *gin.Engine
}

func (r *Router) Run(addr ...string) {
	r.gin.Run(addr...)
}

func NewRouter(si *httpserver.HttpServer) *Router {
	r := gin.New()
	a := r.Group("/actores")
	a.POST("/access", func(ctx *gin.Context) { si.CheckAccess(NewContext(ctx)) })
	a.POST("/new_token", func(ctx *gin.Context) { si.NewToken(NewContext(ctx)) })
	a.DELETE("/remove_token/:id", func(ctx *gin.Context) { si.RemoveToken(NewContext(ctx)) })
	a.GET("/:id", func(ctx *gin.Context) { si.GetActore(NewContext(ctx)) })
	a.GET("/user/:userId", func(ctx *gin.Context) { si.GetUserActores(NewContext(ctx)) })
	a.POST("/add_ip", func(ctx *gin.Context) { si.AddIp(NewContext(ctx)) })
	a.POST("/remove_ip", func(ctx *gin.Context) { si.RemoveIp(NewContext(ctx)) })
	return &Router{gin: r}

}
