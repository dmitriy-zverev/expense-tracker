package cmd

import "github.com/dmitriy-zverev/expense-tracker/internal/expense"

func add(cmd Command) error {
	exp, err := expense.CreateExpenseObj(
		float64(cmd.Amount),
		cmd.Description,
		cmd.Category,
	)
	if err != nil {
		return err
	}

	if err := expense.AddExpense(exp); err != nil {
		return err
	}

	return nil
}
