package budget

import (
	"encoding/json"
	"os"
	"testing"
	"time"
)

func TestGetBudgets(t *testing.T) {
	setupTestData(t)
	defer cleanupTestData(t)

	// Test empty file
	budgets, err := GetBudgets()
	if err != nil {
		t.Errorf("GetBudgets() error = %v", err)
	}
	if len(budgets) != 0 {
		t.Errorf("GetBudgets() should return empty slice for empty file")
	}

	// Add test data
	testBudgets := []Budget{
		{
			Month:    1,
			Year:     2024,
			Category: "Food",
			Limit:    500.0,
		},
		{
			Month:    1,
			Year:     2024,
			Category: "Transport",
			Limit:    200.0,
		},
	}

	data, _ := json.Marshal(testBudgets)
	os.WriteFile("./data/budgets.json", data, 0755)

	budgets, err = GetBudgets()
	if err != nil {
		t.Errorf("GetBudgets() error = %v", err)
	}
	if len(budgets) != 2 {
		t.Errorf("GetBudgets() length = %v, want 2", len(budgets))
	}
}

func TestGetBudget(t *testing.T) {
	setupTestData(t)
	defer cleanupTestData(t)

	// Add test data
	testBudgets := []Budget{
		{
			Month:    1,
			Year:     2024,
			Category: "Food",
			Limit:    500.0,
		},
		{
			Month:    2,
			Year:     2024,
			Category: "Transport",
			Limit:    200.0,
		},
	}

	data, _ := json.Marshal(testBudgets)
	os.WriteFile("./data/budgets.json", data, 0755)

	tests := []struct {
		name     string
		month    int
		category string
		wantErr  bool
		want     Budget
	}{
		{
			name:     "Valid budget - Food",
			month:    1,
			category: "Food",
			wantErr:  false,
			want: Budget{
				Month:    1,
				Year:     2024,
				Category: "Food",
				Limit:    500.0,
			},
		},
		{
			name:     "Valid budget - Transport",
			month:    2,
			category: "Transport",
			wantErr:  false,
			want: Budget{
				Month:    2,
				Year:     2024,
				Category: "Transport",
				Limit:    200.0,
			},
		},
		{
			name:     "Non-existent budget",
			month:    3,
			category: "Entertainment",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			budget, err := GetBudget(tt.month, tt.category)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetBudget() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if budget.Month != tt.want.Month {
					t.Errorf("GetBudget() month = %v, want %v", budget.Month, tt.want.Month)
				}
				if budget.Category != tt.want.Category {
					t.Errorf("GetBudget() category = %v, want %v", budget.Category, tt.want.Category)
				}
				if budget.Limit != tt.want.Limit {
					t.Errorf("GetBudget() limit = %v, want %v", budget.Limit, tt.want.Limit)
				}
			}
		})
	}
}

func TestSetBudget(t *testing.T) {
	setupTestData(t)
	defer cleanupTestData(t)

	tests := []struct {
		name     string
		month    int
		category string
		limit    float64
		wantErr  bool
	}{
		{
			name:     "Valid budget",
			month:    1,
			category: "Food",
			limit:    500.0,
			wantErr:  false,
		},
		{
			name:     "Invalid month - too low",
			month:    0,
			category: "Food",
			limit:    500.0,
			wantErr:  true,
		},
		{
			name:     "Invalid month - too high",
			month:    13,
			category: "Food",
			limit:    500.0,
			wantErr:  true,
		},
		{
			name:     "Empty category",
			month:    1,
			category: "",
			limit:    500.0,
			wantErr:  true,
		},
		{
			name:     "Negative limit",
			month:    1,
			category: "Food",
			limit:    -100.0,
			wantErr:  true,
		},
		{
			name:     "Zero limit",
			month:    1,
			category: "Food",
			limit:    0.0,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := SetBudget(tt.month, tt.category, tt.limit)

			if (err != nil) != tt.wantErr {
				t.Errorf("SetBudget() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				// Verify budget was set
				budget, err := GetBudget(tt.month, tt.category)
				if err != nil {
					t.Errorf("GetBudget() error = %v", err)
				}
				if budget.Limit != tt.limit {
					t.Errorf("Expected limit %v, got %v", tt.limit, budget.Limit)
				}
				if budget.Year != time.Now().Year() {
					t.Errorf("Expected year %v, got %v", time.Now().Year(), budget.Year)
				}
			}
		})
	}
}

func TestSetBudgetUpdate(t *testing.T) {
	setupTestData(t)
	defer cleanupTestData(t)

	// Set initial budget
	err := SetBudget(1, "Food", 500.0)
	if err != nil {
		t.Errorf("SetBudget() error = %v", err)
	}

	// Update the same budget
	err = SetBudget(1, "Food", 600.0)
	if err != nil {
		t.Errorf("SetBudget() update error = %v", err)
	}

	// Verify budget was updated, not duplicated
	budgets, err := GetBudgets()
	if err != nil {
		t.Errorf("GetBudgets() error = %v", err)
	}
	if len(budgets) != 1 {
		t.Errorf("Expected 1 budget, got %d", len(budgets))
	}
	if budgets[0].Limit != 600.0 {
		t.Errorf("Expected limit 600.0, got %v", budgets[0].Limit)
	}
}

func TestRemoveBudget(t *testing.T) {
	setupTestData(t)
	defer cleanupTestData(t)

	// Add test data
	testBudgets := []Budget{
		{
			Month:    1,
			Year:     2024,
			Category: "Food",
			Limit:    500.0,
		},
		{
			Month:    1,
			Year:     2024,
			Category: "Transport",
			Limit:    200.0,
		},
	}

	data, _ := json.Marshal(testBudgets)
	os.WriteFile("./data/budgets.json", data, 0755)

	tests := []struct {
		name     string
		month    int
		category string
		wantErr  bool
	}{
		{
			name:     "Valid removal",
			month:    1,
			category: "Food",
			wantErr:  false,
		},
		{
			name:     "Non-existent budget",
			month:    3,
			category: "Entertainment",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := RemoveBudget(tt.month, tt.category)

			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveBudget() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				// Verify budget was removed
				_, err := GetBudget(tt.month, tt.category)
				if err == nil {
					t.Errorf("Budget should have been removed")
				}
			}
		})
	}
}

func TestGetBudgetLimit(t *testing.T) {
	setupTestData(t)
	defer cleanupTestData(t)

	// Add test data
	testBudgets := []Budget{
		{
			Month:    1,
			Year:     2024,
			Category: "Food",
			Limit:    500.0,
		},
	}

	data, _ := json.Marshal(testBudgets)
	os.WriteFile("./data/budgets.json", data, 0755)

	tests := []struct {
		name     string
		month    int
		category string
		want     float64
		wantErr  bool
	}{
		{
			name:     "Valid budget limit",
			month:    1,
			category: "Food",
			want:     500.0,
			wantErr:  false,
		},
		{
			name:     "Non-existent budget",
			month:    2,
			category: "Transport",
			want:     0.0,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			limit, err := GetBudgetLimit(tt.month, tt.category)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetBudgetLimit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && limit != tt.want {
				t.Errorf("GetBudgetLimit() = %v, want %v", limit, tt.want)
			}
		})
	}
}

func TestGetBudgetLimitsForMonth(t *testing.T) {
	setupTestData(t)
	defer cleanupTestData(t)

	// Add test data
	testBudgets := []Budget{
		{
			Month:    1,
			Year:     2024,
			Category: "Food",
			Limit:    500.0,
		},
		{
			Month:    1,
			Year:     2024,
			Category: "Transport",
			Limit:    200.0,
		},
		{
			Month:    2,
			Year:     2024,
			Category: "Food",
			Limit:    600.0,
		},
	}

	data, _ := json.Marshal(testBudgets)
	os.WriteFile("./data/budgets.json", data, 0755)

	tests := []struct {
		name      string
		month     int
		wantCount int
	}{
		{
			name:      "Month with multiple budgets",
			month:     1,
			wantCount: 2,
		},
		{
			name:      "Month with single budget",
			month:     2,
			wantCount: 1,
		},
		{
			name:      "Month with no budgets",
			month:     3,
			wantCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			budgets, err := GetBudgetLimitsForMonth(tt.month)
			if err != nil {
				t.Errorf("GetBudgetLimitsForMonth() error = %v", err)
				return
			}

			if len(budgets) != tt.wantCount {
				t.Errorf("GetBudgetLimitsForMonth() count = %v, want %v", len(budgets), tt.wantCount)
			}

			// Verify all returned budgets are for the correct month
			for _, budget := range budgets {
				if budget.Month != tt.month {
					t.Errorf("Budget month = %v, want %v", budget.Month, tt.month)
				}
			}
		})
	}
}

func TestValidateBudgetParams(t *testing.T) {
	tests := []struct {
		name     string
		month    int
		category string
		limit    float64
		wantOk   bool
		wantErr  string
	}{
		{
			name:     "Valid parameters",
			month:    1,
			category: "Food",
			limit:    500.0,
			wantOk:   true,
		},
		{
			name:     "Invalid month - too low",
			month:    0,
			category: "Food",
			limit:    500.0,
			wantOk:   false,
			wantErr:  "invalid month",
		},
		{
			name:     "Invalid month - too high",
			month:    13,
			category: "Food",
			limit:    500.0,
			wantOk:   false,
			wantErr:  "invalid month",
		},
		{
			name:     "Empty category",
			month:    1,
			category: "",
			limit:    500.0,
			wantOk:   false,
			wantErr:  "category not set",
		},
		{
			name:     "Negative limit",
			month:    1,
			category: "Food",
			limit:    -100.0,
			wantOk:   false,
			wantErr:  "limit cannot be less then zero",
		},
		{
			name:     "Zero limit",
			month:    1,
			category: "Food",
			limit:    0.0,
			wantOk:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok, err := validateBudgetParams(tt.month, tt.category, tt.limit)

			if ok != tt.wantOk {
				t.Errorf("validateBudgetParams() ok = %v, want %v", ok, tt.wantOk)
			}

			if !tt.wantOk && err.Error() != tt.wantErr {
				t.Errorf("validateBudgetParams() error = %v, want %v", err.Error(), tt.wantErr)
			}
		})
	}
}

func TestIsBudgetAlreadySet(t *testing.T) {
	budgets := []Budget{
		{
			Month:    1,
			Year:     2024,
			Category: "Food",
			Limit:    500.0,
		},
		{
			Month:    2,
			Year:     2024,
			Category: "Transport",
			Limit:    200.0,
		},
	}

	tests := []struct {
		name      string
		budget    Budget
		wantIdx   int
		wantIsSet bool
	}{
		{
			name: "Budget already set",
			budget: Budget{
				Month:    1,
				Year:     2024,
				Category: "Food",
				Limit:    600.0, // Different limit, but same month/year/category
			},
			wantIdx:   0,
			wantIsSet: true,
		},
		{
			name: "Budget not set",
			budget: Budget{
				Month:    3,
				Year:     2024,
				Category: "Entertainment",
				Limit:    300.0,
			},
			wantIdx:   -1,
			wantIsSet: false,
		},
		{
			name: "Different year",
			budget: Budget{
				Month:    1,
				Year:     2023,
				Category: "Food",
				Limit:    500.0,
			},
			wantIdx:   -1,
			wantIsSet: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idx, isSet := isBudgetAlreadySet(budgets, tt.budget)

			if idx != tt.wantIdx {
				t.Errorf("isBudgetAlreadySet() idx = %v, want %v", idx, tt.wantIdx)
			}
			if isSet != tt.wantIsSet {
				t.Errorf("isBudgetAlreadySet() isSet = %v, want %v", isSet, tt.wantIsSet)
			}
		})
	}
}

// Helper functions for test setup
func setupTestData(t *testing.T) {
	// Create test data directory
	err := os.MkdirAll("./data", 0755)
	if err != nil {
		t.Fatalf("Failed to create test data directory: %v", err)
	}
	// Create empty budgets.json file
	err = os.WriteFile("./data/budgets.json", []byte("[]"), 0755)
	if err != nil {
		t.Fatalf("Failed to create test budgets file: %v", err)
	}
}

func cleanupTestData(t *testing.T) {
	err := os.RemoveAll("./data")
	if err != nil {
		t.Logf("Failed to cleanup test data: %v", err)
	}
}
