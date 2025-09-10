package utils

import (
	"testing"
	"time"

	"github.com/dmitriy-zverev/expense-tracker/internal/expense"
)

func TestIsExpenseEmpty(t *testing.T) {
	tests := []struct {
		name    string
		expense expense.Expense
		want    bool
	}{
		{
			name: "Empty expense",
			expense: expense.Expense{
				Description: "",
				Amount:      -1,
				Category:    "",
			},
			want: true,
		},
		{
			name: "Non-empty expense - all fields filled",
			expense: expense.Expense{
				Description: "Test expense",
				Amount:      100.0,
				Category:    "Food",
			},
			want: false,
		},
		{
			name: "Non-empty expense - only description filled",
			expense: expense.Expense{
				Description: "Test expense",
				Amount:      -1,
				Category:    "",
			},
			want: false,
		},
		{
			name: "Non-empty expense - only amount filled",
			expense: expense.Expense{
				Description: "",
				Amount:      100.0,
				Category:    "",
			},
			want: false,
		},
		{
			name: "Non-empty expense - only category filled",
			expense: expense.Expense{
				Description: "",
				Amount:      -1,
				Category:    "Food",
			},
			want: false,
		},
		{
			name: "Non-empty expense - zero amount",
			expense: expense.Expense{
				Description: "Free item",
				Amount:      0,
				Category:    "Gift",
			},
			want: false,
		},
		{
			name: "Edge case - amount exactly -1",
			expense: expense.Expense{
				Description: "Test",
				Amount:      -1,
				Category:    "Test",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsExpenseEmpty(tt.expense)
			if result != tt.want {
				t.Errorf("IsExpenseEmpty() = %v, want %v", result, tt.want)
			}
		})
	}
}

func TestIsExpenseAmountValid(t *testing.T) {
	tests := []struct {
		name    string
		expense expense.Expense
		want    bool
	}{
		{
			name: "Valid positive amount",
			expense: expense.Expense{
				Amount: 100.50,
			},
			want: true,
		},
		{
			name: "Valid zero amount",
			expense: expense.Expense{
				Amount: 0,
			},
			want: true,
		},
		{
			name: "Invalid negative amount",
			expense: expense.Expense{
				Amount: -50.0,
			},
			want: false,
		},
		{
			name: "Invalid amount -1",
			expense: expense.Expense{
				Amount: -1,
			},
			want: false,
		},
		{
			name: "Very small positive amount",
			expense: expense.Expense{
				Amount: 0.01,
			},
			want: true,
		},
		{
			name: "Large amount",
			expense: expense.Expense{
				Amount: 999999.99,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsExpenseAmountValid(tt.expense)
			if result != tt.want {
				t.Errorf("IsExpenseAmountValid() = %v, want %v", result, tt.want)
			}
		})
	}
}

func TestIsExpenseValid(t *testing.T) {
	now := time.Now().UTC()

	tests := []struct {
		name    string
		expense expense.Expense
		want    bool
	}{
		{
			name: "Valid complete expense",
			expense: expense.Expense{
				ID:          1,
				Amount:      100.50,
				Description: "Lunch",
				Category:    "Food",
				Date:        now,
				Month:       int(now.Month()),
				IsDeleted:   false,
			},
			want: true,
		},
		{
			name: "Valid expense with zero amount",
			expense: expense.Expense{
				ID:          2,
				Amount:      0,
				Description: "Free sample",
				Category:    "Gift",
				Date:        now,
				Month:       int(now.Month()),
				IsDeleted:   false,
			},
			want: true,
		},
		{
			name: "Invalid - empty expense",
			expense: expense.Expense{
				Description: "",
				Amount:      -1,
				Category:    "",
			},
			want: false,
		},
		{
			name: "Invalid - negative amount",
			expense: expense.Expense{
				ID:          3,
				Amount:      -50.0,
				Description: "Invalid expense",
				Category:    "Test",
				Date:        now,
				Month:       int(now.Month()),
				IsDeleted:   false,
			},
			want: false,
		},
		{
			name: "Valid - empty but positive amount",
			expense: expense.Expense{
				Description: "",
				Amount:      100.0,
				Category:    "",
			},
			want: true, // Valid because amount is valid and not empty (amount != -1)
		},
		{
			name: "Valid - minimal valid expense",
			expense: expense.Expense{
				Amount:      1.0,
				Description: "Test",
				Category:    "Test",
			},
			want: true,
		},
		{
			name: "Valid - only description different from empty criteria",
			expense: expense.Expense{
				Amount:      0,
				Description: "Free item",
				Category:    "",
			},
			want: true, // Valid because amount is valid and not empty (amount != -1)
		},
		{
			name: "Edge case - amount -1 with other fields filled",
			expense: expense.Expense{
				Amount:      -1,
				Description: "Test",
				Category:    "Test",
			},
			want: false, // Invalid amount
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsExpenseValid(tt.expense)
			if result != tt.want {
				t.Errorf("IsExpenseValid() = %v, want %v", result, tt.want)

				// Additional debugging info
				isEmpty := IsExpenseEmpty(tt.expense)
				isAmountValid := IsExpenseAmountValid(tt.expense)
				t.Logf("  IsExpenseEmpty: %v", isEmpty)
				t.Logf("  IsExpenseAmountValid: %v", isAmountValid)
				t.Logf("  Expected: %v", tt.want)
			}
		})
	}
}

func TestValidationFunctionsIntegration(t *testing.T) {
	// Test the relationship between the validation functions
	t.Run("Integration test - validation function relationships", func(t *testing.T) {
		testCases := []struct {
			name              string
			expense           expense.Expense
			expectEmpty       bool
			expectAmountValid bool
			expectValid       bool
		}{
			{
				name: "Perfect expense",
				expense: expense.Expense{
					Amount:      100.0,
					Description: "Test",
					Category:    "Food",
				},
				expectEmpty:       false,
				expectAmountValid: true,
				expectValid:       true,
			},
			{
				name: "Empty expense",
				expense: expense.Expense{
					Amount:      -1,
					Description: "",
					Category:    "",
				},
				expectEmpty:       true,
				expectAmountValid: false,
				expectValid:       false,
			},
			{
				name: "Valid amount but empty fields",
				expense: expense.Expense{
					Amount:      100.0,
					Description: "",
					Category:    "",
				},
				expectEmpty:       false, // Not empty because amount != -1
				expectAmountValid: true,
				expectValid:       true, // Valid because amount is valid and not empty
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				isEmpty := IsExpenseEmpty(tc.expense)
				isAmountValid := IsExpenseAmountValid(tc.expense)
				isValid := IsExpenseValid(tc.expense)

				if isEmpty != tc.expectEmpty {
					t.Errorf("IsExpenseEmpty() = %v, want %v", isEmpty, tc.expectEmpty)
				}
				if isAmountValid != tc.expectAmountValid {
					t.Errorf("IsExpenseAmountValid() = %v, want %v", isAmountValid, tc.expectAmountValid)
				}
				if isValid != tc.expectValid {
					t.Errorf("IsExpenseValid() = %v, want %v", isValid, tc.expectValid)
				}

				// Verify the relationship: IsExpenseValid should be true only if
				// both IsExpenseAmountValid is true AND IsExpenseEmpty is false
				expectedValid := isAmountValid && !isEmpty
				if isValid != expectedValid {
					t.Errorf("IsExpenseValid() relationship broken: got %v, expected %v (amountValid: %v, empty: %v)",
						isValid, expectedValid, isAmountValid, isEmpty)
				}
			})
		}
	})
}

// Benchmark tests to ensure validation functions are performant
func BenchmarkIsExpenseEmpty(b *testing.B) {
	exp := expense.Expense{
		Amount:      100.0,
		Description: "Test expense",
		Category:    "Food",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsExpenseEmpty(exp)
	}
}

func BenchmarkIsExpenseAmountValid(b *testing.B) {
	exp := expense.Expense{
		Amount: 100.0,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsExpenseAmountValid(exp)
	}
}

func BenchmarkIsExpenseValid(b *testing.B) {
	exp := expense.Expense{
		Amount:      100.0,
		Description: "Test expense",
		Category:    "Food",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsExpenseValid(exp)
	}
}
