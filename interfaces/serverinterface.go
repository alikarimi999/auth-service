package interfaces

type ServerInterface interface {
	CheckAccess(ServerContext)
	AddActore(ServerContext)
	GetActore(ServerContext)
}
