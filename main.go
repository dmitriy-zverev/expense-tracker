package main

import (
	"fmt"
	"os"

	"github.com/dmitriy-zverev/expense-tracker/cmd"
)

func main() {
	command, err := cmd.ParseCommand(os.Args)
	if err != nil {
		fmt.Printf("\nError: %v\n", err)
		return
	}

	if err := command.Run(); err != nil {
		fmt.Printf("\nError: %v\n", err)
		os.Exit(1)
	}
}
