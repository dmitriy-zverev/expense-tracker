package cmd

import (
	"errors"

	"github.com/dmitriy-zverev/expense-tracker/internal/expense"
)

/**
* Deletes an expense by its ID.
*
* @param cmd The command containing the ID of the expense to delete.
* @return An error if the ID is not provided or if the deletion fails; otherwise, nil.
 */
func delete(cmd Command) error {
	if cmd.ID == -1 {
		return errors.New("id not provided")
	}

	if err := expense.DeleteExpense(cmd.ID); err != nil {
		return err
	}

	return nil
}
