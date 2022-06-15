package httpserver

import (
	"github.com/billsbook/auth/app"
	"github.com/billsbook/auth/interfaces"
)

type httpServer struct {
	App *app.Application
}

func NewServer(app *app.Application) interfaces.ServerInterface {
	return &httpServer{App: app}
}
