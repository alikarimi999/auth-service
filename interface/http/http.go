package httpserver

import (
	"github.com/alikarimi999/auth_service/app"
)

type HttpServer struct {
	App *app.Application
}

func NewServer(app *app.Application) *HttpServer {
	return &HttpServer{App: app}
}
