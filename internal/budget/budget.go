package budget

import (
	"encoding/json"
	"errors"
	"slices"
	"time"

	"github.com/dmitriy-zverev/expense-tracker/internal/storage"
)

type Budget struct {
	Month    int     `json:"month"`
	Year     int     `json:"year"`
	Category string  `json:"category"`
	Limit    float64 `json:"limit"`
}

const (
	DEFAULT_BUDGET_FILE_PATH = "./data/budgets.json"
)

func GetBudgets() ([]Budget, error) {
	data, err := storage.GetFileData(DEFAULT_BUDGET_FILE_PATH)
	if err != nil {
		return []Budget{}, err
	}

	if len(data) < 1 {
		return []Budget{}, nil
	}

	var budgets []Budget
	if err := json.Unmarshal(data, &budgets); err != nil {
		return []Budget{}, err
	}

	return budgets, nil
}

func GetBudget(month int, category string) (Budget, error) {
	budgets, err := GetBudgets()
	if err != nil {
		return Budget{}, err
	}

	for _, budget := range budgets {
		if budget.Month == month && budget.Category == category {
			return budget, nil
		}
	}

	return Budget{}, errors.New("budget not found")
}

func SetBudget(month int, category string, limit float64) error {
	if ok, err := validateBudgetParams(month, category, limit); !ok {
		return err
	}

	budgets, err := GetBudgets()
	if err != nil {
		return err
	}

	budget := Budget{
		Month:    month,
		Year:     time.Now().Year(),
		Category: category,
		Limit:    float64(limit),
	}

	idx, ok := isBudgetAlreadySet(budgets, budget)
	if !ok {
		budgets = append(budgets, budget)
	} else {
		budgets[idx] = budget
	}

	data, err := json.Marshal(budgets)
	if err != nil {
		return err
	}

	if err := storage.WriteFileData(DEFAULT_BUDGET_FILE_PATH, data); err != nil {
		return err
	}

	return nil
}

func RemoveBudget(month int, category string) error {
	budgets, err := GetBudgets()
	if err != nil {
		return err
	}

	budget, err := GetBudget(month, category)
	if err != nil {
		return err
	}

	idx := slices.Index(budgets, budget)
	budgets = append(budgets[:idx], budgets[idx+1:]...)

	data, err := json.Marshal(budgets)
	if err != nil {
		return err
	}

	if err := storage.WriteFileData(DEFAULT_BUDGET_FILE_PATH, data); err != nil {
		return err
	}

	return nil
}

func GetBudgetLimit(month int, category string) (float64, error) {
	budget, err := GetBudget(month, category)
	if err != nil {
		return 0.0, err
	}

	return budget.Limit, nil
}

func GetBudgetLimitsForMonth(month int) ([]Budget, error) {
	budgets, err := GetBudgets()
	if err != nil {
		return []Budget{}, err
	}

	resultBudgets := []Budget{}

	for _, b := range budgets {
		if b.Month == month {
			resultBudgets = append(resultBudgets, b)
		}
	}

	return resultBudgets, nil
}

func validateBudgetParams(month int, category string, limit float64) (bool, error) {
	if month < 1 || month > 12 {
		return false, errors.New("invalid month")
	}

	if category == "" {
		return false, errors.New("category not set")
	}

	if limit < 0 {
		return false, errors.New("limit cannot be less then zero")
	}

	return true, nil
}

func isBudgetAlreadySet(budgets []Budget, budget Budget) (int, bool) {
	isSet := false
	idx := -1

	for i, b := range budgets {
		if b.Month == budget.Month && b.Year == budget.Year && b.Category == budget.Category {
			isSet = true
			idx = i
			break
		}
	}

	return idx, isSet
}
