package cmd

import (
	"github.com/dmitriy-zverev/expense-tracker/internal/expense"
)

func update(cmd Command) error {
	exp, err := expense.GetExpense(cmd.ID)
	if err != nil {
		return err
	}

	newDesc := cmd.Description
	newAmount := float64(cmd.Amount)
	newCategory := cmd.Category

	if newDesc == "" && exp.Description != "" {
		newDesc = exp.Description
	}

	if newAmount == -1 && exp.Amount != -1 {
		newAmount = exp.Amount
	}

	if newCategory == "" && exp.Category != "" {
		newCategory = exp.Category
	}

	if err := expense.UpdateExpense(cmd.ID, newAmount, newDesc, newCategory); err != nil {
		return err
	}

	return nil
}
