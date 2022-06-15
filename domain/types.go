package domain

type Resource int

const (
	Products Resource = iota
	Plans
	Subscriptions
)

func (r Resource) String() string {
	switch r {
	case Products:
		return "products"
	case Plans:
		return "plans"
	case Subscriptions:
		return "subscriptions"
	default:
		return ""
	}
}

type Action int

const (
	None Action = iota
	Read
	Write
)

func (a Action) String() string {
	switch a {
	case Read:
		return "read"
	case Write:
		return "write"
	default:
		return "none"
	}
}

type ActoreType int

const (
	Token ActoreType = iota // APIToken struct
	Acc                     // Account struct
)

type Permission struct {
	Resource Resource
	Action   Action
}

type Tenat struct {
	ID string
}
