package dto

import "github.com/alikarimi999/auth_service/domain"

func Resource(res string) domain.Resource {
	switch res {
	case "orders":
		return domain.Orders
	default:
		return -1
	}
}

func Action(act string) domain.Action {
	switch act {
	case "read":
		return domain.Read
	case "write":
		return domain.Write
	default:
		return domain.None
	}
}
