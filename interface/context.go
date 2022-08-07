package interfaces

import "net/http"

type ServerContext interface {
	Param(string) string
	Bind(interface{}) error
	JSON(int, interface{})
	GetKey(key string) (interface{}, bool)
	SetKey(key string, value interface{})
	Request() *http.Request
}
