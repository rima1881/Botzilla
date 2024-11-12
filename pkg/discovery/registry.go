package discovery

// Each component can be a member of multiple groups
type Component struct {
	ID    string
	group []string
}

// Keys are the IDs
var components = make(map[string]Component)

func Register(comp Component) {
	components[comp.ID] = comp
}

func GetComponents() {

}
