package expense

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/dmitriy-zverev/expense-tracker/internal/storage"
)

type Expense struct {
	Amount      float64   `json:"amount"`
	Date        time.Time `json:"date"`
	ID          int       `json:"id"`
	Month       int       `json:"month"`
	IsDeleted   bool      `json:"is_deleted"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
}

const (
	EXPENSES_FILE_PATH = "./data/expenses.json"
)

func CreateExpenseObj(amount float64, desc, category string) (Expense, error) {
	expenses, err := GetExpenses()
	if err != nil {
		return Expense{}, err
	}

	date := time.Now().UTC()

	return Expense{
		Amount:      amount,
		Description: desc,
		Category:    category,
		IsDeleted:   false,
		Date:        date,
		ID:          len(expenses),
		Month:       int(date.Month()),
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
	expenses, err := GetExpenses()
	if err != nil {
		return Expense{}, err
	}

	if id > len(expenses) || id < 0 {
		return Expense{}, errors.New("cannot find expense with provided id")
	}

	return expenses[id], nil
}

func GetExpenseForCategory(category string) (float64, error) {
	expenses, err := GetExpenses()
	if err != nil {
		return 0.0, err
	}

	totalExpenses := 0.0

	for _, e := range expenses {
		if e.Category == category {
			totalExpenses += e.Amount
		}
	}

	return totalExpenses, nil
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

	if id > len(expenses) || id < 0 {
		return errors.New("cannot find expense with provided id")
	}

	expenses[id].IsDeleted = true

	fmt.Println(expenses)

	data, err := json.Marshal(expenses)
	if err != nil {
		return err
	}

	fmt.Println(string(data))

	if err := storage.WriteFileData(EXPENSES_FILE_PATH, data); err != nil {
		return err
	}

	return nil
}

func UpdateExpense(id int, amount float64, desc, category string) error {
	expenses, err := GetExpenses()
	if err != nil {
		return err
	}

	if id > len(expenses) || id < 0 {
		return errors.New("cannot find expense with provided id")
	}

	expenses[id].Amount = float64(amount)
	expenses[id].Description = desc
	expenses[id].Category = category

	data, err := json.Marshal(expenses)
	if err != nil {
		return err
	}

	if err := storage.WriteFileData(EXPENSES_FILE_PATH, data); err != nil {
		return err
	}

	return nil
}
