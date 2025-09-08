package cmd

type cliCommand struct {
	Name        string
	Description string
	Callback    func(Command) error
}

var commands map[string]cliCommand

func initCommands() {
	commands = map[string]cliCommand{
		"add": {
			Name:        "add",
			Description: "Adds expense to your tracker",
			Callback:    add,
		},
	}
}
