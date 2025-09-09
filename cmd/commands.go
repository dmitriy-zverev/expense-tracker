package cmd

type cliCommand struct {
	Name        string
	Description string
	Callback    func(Command) error
}

var commands map[string]cliCommand

/**
* Initializes the commands map with available CLI commands.
* Each command is defined with a name, description, and callback function.
* Supported commands:
* - "add": Adds an expense to the tracker
* - "list": Lists all expenses
* - "delete": Deletes an expense by its ID
* - "update": Updates an expense by its id
* - "summary": Summarizes all expenses—if set within provided month
 */
func initCommands() {
	commands = map[string]cliCommand{
		"add": {
			Name:        "add",
			Description: "Adds expense to your tracker",
			Callback:    add,
		},
		"list": {
			Name:        "list",
			Description: "Lists all of the expenses",
			Callback:    list,
		},
		"delete": {
			Name:        "delete",
			Description: "Deletes expense with provided id",
			Callback:    delete,
		},
		"update": {
			Name:        "update",
			Description: "Updates expense with provided id",
			Callback:    update,
		},
		"summary": {
			Name:        "summary",
			Description: "Summarizes all expenses—if set within provided month",
			Callback:    summary,
		},
		"export": {
			Name:        "export",
			Description: "Exports expenses into a .csv file—if set with custom file name",
			Callback:    export,
		},
	}
}
