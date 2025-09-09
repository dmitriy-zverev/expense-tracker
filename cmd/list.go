package cmd

import (
	"fmt"
	"strings"

	"github.com/dmitriy-zverev/expense-tracker/internal/expense"
)

/**
* Lists all expenses, optionally including deleted ones based on the command configuration.
*
* @param cmd The command object containing configuration options, including whether to include deleted expenses.
* @return An error if there was an issue retrieving the expenses; otherwise, nil.
 */
func list(cmd Command) error {
	expenses, err := expense.GetExpenses()
	if err != nil {
		return err
	}

	fmt.Printf(
		"# ID\tDate\t\tDescription%sAmount\tCategory\n",
		strings.Repeat(" ", PRINT_MAX_DESCRIPTION_LENGTH-len("Description")+1),
	)

	for _, exp := range expenses {
		if !cmd.WithDeleted && exp.IsDeleted {
			continue
		}

		if cmd.Month != -1 && exp.Month != cmd.Month {
			continue
		}

		year, month, day := exp.Date.Date()
		dateString := fmt.Sprintf("%d-%d-%d", year, month, day)

		maxLen := min(PRINT_MAX_DESCRIPTION_LENGTH, len(exp.Description))

		spaces := strings.Repeat(" ", PRINT_MAX_DESCRIPTION_LENGTH-maxLen+1)
		fmt.Printf(
			"# %d\t%s\t%s%s%.2f\t%s",
			exp.ID,
			dateString,
			exp.Description[:maxLen],
			spaces,
			exp.Amount,
			exp.Category,
		)

		if exp.IsDeleted {
			fmt.Printf("\t(deleted)")
		}

		fmt.Printf("\n")
	}

	fmt.Println()

	return nil
}
