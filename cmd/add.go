package cmd

import (
	"errors"

	"github.com/dmitriy-zverev/expense-tracker/internal/expense"
	"github.com/dmitriy-zverev/expense-tracker/internal/utils"
)

/**
* Adds a new expense after validating it.
*
* @param cmd The command containing the expense details.
* @return An error if the expense creation, validation, or addition fails; otherwise, nil.
 */
func add(cmd Command) error {
	exp, err := expense.CreateExpenseObj(
		float64(cmd.Amount),
		cmd.Description,
		cmd.Category,
	)
	if err != nil {
		return err
	}

	if !utils.IsExpenseValid(exp) {
		return errors.New("not valid expense")
	}

	if err := expense.AddExpense(exp); err != nil {
		return err
	}

	return nil
}
