package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dmitriy-zverev/expense-tracker/internal/budget"
)

const (
	CATEGORY_LIMIT_CHARS = 20
)

func budgetCmd(cmd Command) error {
	switch cmd.BudgetCmd {
	case BUDGET_SET_CMD:
		if err := setBudget(cmd); err != nil {
			return err
		}
	case BUDGET_LIST_CMD:
		if err := listBudget(); err != nil {
			return err
		}
	case BUDGET_REMOVE_CMD:
		if err := removeBudget(cmd); err != nil {
			return err
		}
	default:
		return errors.New("command for budget is not provided")
	}

	return nil
}

func setBudget(cmd Command) error {
	if err := budget.SetBudget(cmd.Month, cmd.Category, cmd.Limit); err != nil {
		return err
	}

	return nil
}

func listBudget() error {
	budgets, err := budget.GetBudgets()
	if err != nil {
		return err
	}

	fmt.Printf(
		"#\tMonth\tYear\tCategory%sLimit\n",
		strings.Repeat(" ", CATEGORY_LIMIT_CHARS-len("category")+1),
	)

	for _, budget := range budgets {
		categoryStringLen := min(CATEGORY_LIMIT_CHARS, len(budget.Category))

		fmt.Printf(
			"#\t%d\t%d\t%s%s%.2f\n",
			budget.Month,
			budget.Year,
			budget.Category[:categoryStringLen],
			strings.Repeat(" ", CATEGORY_LIMIT_CHARS-categoryStringLen+1),
			budget.Limit,
		)
	}

	return nil
}

func removeBudget(cmd Command) error {
	if err := budget.RemoveBudget(cmd.Month, cmd.Category); err != nil {
		return err
	}

	return nil
}
