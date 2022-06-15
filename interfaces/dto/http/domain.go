package dto

import "github.com/billsbook/auth/domain"

func ActoreID(id string) domain.ActoreID {
	switch id[:8] {
	case "api_key_":
		return &domain.TokenID{ApiKey: id[8:]}
	default:
		return &domain.AccountID{User: id}
	}
}

func Resource(res string) domain.Resource {
	switch res {
	case "plans":
		return domain.Plans
	case "products":
		return domain.Products
	case "subscriptions":
		return domain.Subscriptions
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
