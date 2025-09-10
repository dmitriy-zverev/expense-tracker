package cmd

import (
	"fmt"
	"time"

	"github.com/dmitriy-zverev/expense-tracker/internal/expense"
)

func summary(cmd Command) error {
	expenses, err := expense.GetExpenses()
	if err != nil {
		return err
	}

	fmt.Printf("Total expenses")

	if cmd.Month != -1 {
		fmt.Printf(" in %v: $", time.Month(cmd.Month).String())
	} else {
		fmt.Printf(": $")
	}

	totalExpenses := 0.0
	for _, exp := range expenses {
		if cmd.Month != -1 && cmd.Month != exp.Month {
			continue
		}

		if cmd.Category != "" && cmd.Category != exp.Category {
			continue
		}

		totalExpenses += exp.Amount
	}

	fmt.Printf("%.2f\n", totalExpenses)

	return nil
}
