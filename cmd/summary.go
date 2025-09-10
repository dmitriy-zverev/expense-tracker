package cmd

import (
	"fmt"
	"time"

	"github.com/dmitriy-zverev/expense-tracker/internal/budget"
	"github.com/dmitriy-zverev/expense-tracker/internal/expense"
)

func summary(cmd Command) error {
	expenses, err := expense.GetExpenses()
	if err != nil {
		return err
	}

	fmt.Printf("Total expenses")

	if cmd.Month != -1 {
		fmt.Printf(" in %v: ", time.Month(cmd.Month).String())
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

	fmt.Printf("%.2f $\n\n", totalExpenses)

	if err := printBudget(cmd, totalExpenses); err != nil {
		return err
	}

	return nil
}

func printBudget(cmd Command, expenses float64) error {
	if cmd.Month != -1 && cmd.Category != "" {
		budgetLimit, err := budget.GetBudgetLimit(cmd.Month, cmd.Category)
		if err != nil {
			return err
		}

		fmt.Printf(
			"Budget for '%s' in %v: %.2f $\n",
			cmd.Category,
			time.Month(cmd.Month).String(),
			budgetLimit,
		)

		fmt.Printf("Current budget stat: %.2f $\n", budgetLimit-expenses)
	}

	if cmd.Month == -1 && cmd.Category != "" {
		budgetLimit, err := budget.GetBudgetLimit(int(time.Now().Month()), cmd.Category)
		if err != nil {
			return err
		}

		fmt.Printf(
			"Budget for '%s' in %v: %.2f $\n",
			cmd.Category,
			time.Month(time.Now().Month()).String(),
			budgetLimit,
		)

		fmt.Printf("Current budget stat: %.2f $\n", budgetLimit-expenses)
	}

	if cmd.Month == -1 && cmd.Category == "" {
		budgets, err := budget.GetBudgetLimitsForMonth(int(time.Now().Month()))
		if err != nil {
			return err
		}

		if len(budgets) < 1 {
			return nil
		}

		for _, b := range budgets {
			fmt.Printf(
				"	Budget for '%s' in %v: %.2f $\n",
				b.Category,
				time.Month(time.Now().Month()).String(),
				b.Limit,
			)

			categoryExpense, err := expense.GetExpenseForCategory(b.Category)
			if err != nil {
				return err
			}

			fmt.Printf("	Current budgeting: %.2f $\n\n", b.Limit-categoryExpense)
		}
	}

	return nil
}
