package cmd

import (
	"fmt"
	"os"

	"github.com/dmitriy-zverev/expense-tracker/internal/expense"
)

func export(cmd Command) error {
	expenses, err := expense.GetExpenses()
	if err != nil {
		return err
	}

	exportFilePath := DEFAULT_EXPORT_FILE_PATH
	if cmd.Output != "" {
		exportFilePath = "./csv/" + cmd.Output
	}

	if _, err := os.Stat("./csv"); os.IsNotExist(err) {
		if err := os.Mkdir("./csv", 0755); err != nil {
			return err
		}
	}

	exportFile, err := os.OpenFile(exportFilePath, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		return err
	}
	defer exportFile.Close()

	csvString := "ID,Date,Description,Amount,Category,Month\n"
	for _, exp := range expenses {
		csvString += fmt.Sprintf(
			"%d,%v,%s,%.2f,%s,%d\n",
			exp.ID,
			exp.Date.String(),
			exp.Description,
			exp.Amount,
			exp.Category,
			exp.Month,
		)
	}

	if _, err := exportFile.WriteString(csvString); err != nil {
		return err
	}

	fmt.Printf("Data has been successfully exported at '%s'\n", exportFilePath)

	return nil
}
