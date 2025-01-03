package core

type Component struct {
	Name    string
	Group   []string
	Address string
	token   string
}

func newComponent(name string, address string) *Component {
	return &Component{
		Name:    name,
		Group:   nil,
		Address: address,
		token:   "test123",
	}
}

func (c *Component) GetToken() string {
	return "test123"
}
