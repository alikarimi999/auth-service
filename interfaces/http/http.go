package httpserver

import (
	"github.com/billsbook/auth_service/app"
	"github.com/billsbook/auth_service/interfaces"
)

type httpServer struct {
	App *app.Application
}

func NewServer(app *app.Application) interfaces.ServerInterface {
	return &httpServer{App: app}
}
