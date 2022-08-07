package domain

type Resource int

const (
	Orders Resource = iota
)

func (r Resource) String() string {
	switch r {
	case Orders:
		return "orders"
	default:
		return ""
	}
}

func ParseResource(s string) Resource {
	switch s {
	case "orders":
		return Orders
	default:
		return -1
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

func ParseAction(s string) Action {
	switch s {
	case "read":
		return Read
	case "write":
		return Write
	default:
		return None
	}
}

type ActoreType int

const (
	Token   ActoreType = iota // APIToken struct
	Account                   // Account struct
)

type Permission struct {
	Resource Resource
	Action   Action
}
