package expense

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/dmitriy-zverev/expense-tracker/internal/storage"
)

type Expense struct {
	Amount      float64   `json:"amount"`
	Date        time.Time `json:"date"`
	ID          int       `json:"id"`
	IsDeleted   bool      `json:"is_deleted"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
}

const (
	EXPENSES_FILE_PATH = "./data/expenses.json"
)

func rewriteExpenses(expenses []Expense) error {
	data, err := json.Marshal(expenses)
	if err != nil {
		return err
	}

	if err := storage.WriteFileData(EXPENSES_FILE_PATH, data); err != nil {
		return err
	}

	return nil
}

func CreateExpenseObj(amount float64, desc, category string) (Expense, error) {
	expenses, err := GetExpenses()
	if err != nil {
		return Expense{}, err
	}

	return Expense{
		Amount:      amount,
		Description: desc,
		Category:    category,
		IsDeleted:   false,
		Date:        time.Now().UTC(),
		ID:          len(expenses),
	}, nil
}

func GetExpenses() ([]Expense, error) {
	fileData, err := storage.GetFileData(EXPENSES_FILE_PATH)
	if err != nil {
		return []Expense{}, err
	}

	expenses := []Expense{}
	if len(fileData) < 1 {
		return []Expense{}, nil
	}

	if err := json.Unmarshal(fileData, &expenses); err != nil {
		return []Expense{}, err
	}

	return expenses, nil
}

func GetExpense(id int) (Expense, error) {
	fileData, err := storage.GetFileData(EXPENSES_FILE_PATH)
	if err != nil {
		return Expense{}, err
	}

	expenses := []Expense{}
	if err := json.Unmarshal(fileData, &expenses); err != nil {
		return Expense{}, err
	}

	if id > len(expenses) {
		return Expense{}, errors.New("cannot find expense with provided id")
	}

	return expenses[id], nil
}

func AddExpense(exp Expense) error {
	expenses, err := GetExpenses()
	if err != nil {
		return err
	}

	expenses = append(expenses, exp)

	data, err := json.Marshal(expenses)
	if err != nil {
		return err
	}

	if err := storage.WriteFileData(EXPENSES_FILE_PATH, data); err != nil {
		return err
	}

	return nil
}

func DeleteExpense(id int) error {
	expenses, err := GetExpenses()
	if err != nil {
		return err
	}

	if id > len(expenses) {
		return errors.New("cannot find expense with provided id")
	}

	expenses[id].IsDeleted = true

	if err := rewriteExpenses(expenses); err != nil {
		return err
	}

	return nil
}

func UpdateExpense(id int, exp Expense) error {
	expenses, err := GetExpenses()
	if err != nil {
		return err
	}

	if id > len(expenses) {
		return errors.New("cannot find expense with provided id")
	}

	expenses[id] = exp

	if err := rewriteExpenses(expenses); err != nil {
		return err
	}

	return nil
}
