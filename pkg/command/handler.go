package command

type CommandHandler interface {
	Handler(data []byte) error
}

type CommandRouter struct {
	handlers map[string]CommandHandler
}

func NewCommandRouter() *CommandRouter {
	return &CommandRouter{
		handlers: map[string]CommandHandler{
			"connect": &ConnectHandler{},
		},
	}
}
