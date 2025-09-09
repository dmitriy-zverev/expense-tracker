package cmd

import (
	"errors"

	"github.com/dmitriy-zverev/expense-tracker/internal/expense"
)

func delete(cmd Command) error {
	if cmd.ID == -1 {
		return errors.New("id not provided")
	}

	if err := expense.DeleteExpense(cmd.ID); err != nil {
		return err
	}

	return nil
}
