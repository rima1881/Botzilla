package core

type Component struct {
	Name       string
	TCPAddress string
	group      []string
}

func NewComponent(name string, address string) *Component {
	return &Component{
		Name:       name,
		TCPAddress: address,
		group:      []string{},
	}
}
