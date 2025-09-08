package main

import (
	"fmt"
	"os"

	"github.com/dmitriy-zverev/expense-tracker/cmd"
)

func main() {
	cmd, err := cmd.ParseCommand(os.Args)
	if err != nil {
		fmt.Printf("\nError: %v\n", err)
		return
	}

	fmt.Println(cmd)
}
