package utils

import "github.com/dmitriy-zverev/expense-tracker/internal/expense"

func IsExpenseEmpty(exp expense.Expense) bool {
	if exp.Description == "" &&
		exp.Amount == -1 &&
		exp.Category == "" {
		return true
	}

	return false
}

func IsExpenseAmountValid(exp expense.Expense) bool {
	return exp.Amount >= 0
}

func IsExpenseValid(exp expense.Expense) bool {
	return IsExpenseAmountValid(exp) && !IsExpenseEmpty(exp)
}
