# Expense Tracker CLI

## Project Overview
Build a command-line expense tracker application in Go that allows users to manage their personal finances with features for adding, updating, deleting, and viewing expenses.

## Technology Stack
- **Language**: Go (based on existing .gitignore)
- **CLI Framework**: cobra (popular Go CLI framework)
- **Data Storage**: JSON file for simplicity and human readability
- **Testing**: Go's built-in testing framework
- **Build Tool**: Go modules

## Project Structure
```
expense-tracker/
├── cmd/
│   ├── root.go          # Root command setup
│   ├── add.go           # Add expense command
│   ├── update.go        # Update expense command
│   ├── delete.go        # Delete expense command
│   ├── list.go          # List expenses command
│   ├── summary.go       # Summary command
│   └── export.go        # Export to CSV command
├── internal/
│   ├── models/
│   │   └── expense.go   # Expense data structure
│   ├── storage/
│   │   └── file.go      # File-based storage operations
│   ├── budget/
│   │   └── budget.go    # Budget management
│   └── utils/
│       └── validation.go # Input validation utilities
├── data/
│   ├── expenses.json    # Main expense data file
│   └── budgets.json     # Budget configuration file
├── main.go              # Application entry point
├── go.mod               # Go module file
├── go.sum               # Go dependencies
└── README.md            # Updated documentation
```

## Data Models

### Expense Structure
```go
type Expense struct {
    ID          int       `json:"id"`
    Date        time.Time `json:"date"`
    Description string    `json:"description"`
    Amount      float64   `json:"amount"`
    Category    string    `json:"category,omitempty"`
}
```

### Budget Structure
```go
type Budget struct {
    Month    int     `json:"month"`
    Year     int     `json:"year"`
    Category string  `json:"category"`
    Limit    float64 `json:"limit"`
}
```

## CLI Commands Design

### Core Commands
1. **add** - Add new expense
   ```bash
   expense-tracker add --description "Lunch" --amount 20 [--category "Food"]
   ```

2. **update** - Update existing expense
   ```bash
   expense-tracker update --id 1 --description "Business Lunch" --amount 25
   ```

3. **delete** - Delete expense
   ```bash
   expense-tracker delete --id 2
   ```

4. **list** - List all expenses
   ```bash
   expense-tracker list [--category "Food"] [--month 8]
   ```

5. **summary** - Show expense summary
   ```bash
   expense-tracker summary [--month 8] [--category "Food"]
   ```

### Additional Commands
6. **export** - Export to CSV
   ```bash
   expense-tracker export --output expenses.csv [--month 8]
   ```

7. **budget** - Manage budgets
   ```bash
   expense-tracker budget set --month 8 --category "Food" --limit 500
   expense-tracker budget list
   ```

## Implementation Phases

### Phase 1: Core Infrastructure
- [ ] Initialize Go module
- [ ] Set up Cobra CLI framework
- [ ] Create basic project structure
- [ ] Implement expense data model
- [ ] Create file-based storage system
- [ ] Implement basic error handling

### Phase 2: Core Commands
- [ ] Implement `add` command
- [ ] Implement `list` command
- [ ] Implement `delete` command
- [ ] Implement `summary` command
- [ ] Add input validation
- [ ] Create comprehensive tests

### Phase 3: Advanced Features
- [ ] Implement `update` command
- [ ] Add category support
- [ ] Implement month-specific filtering
- [ ] Add budget management
- [ ] Implement budget warnings

### Phase 4: Additional Features
- [ ] Implement CSV export
- [ ] Add advanced filtering options
- [ ] Improve error messages
- [ ] Add configuration file support
- [ ] Performance optimizations

### Phase 5: Polish & Documentation
- [ ] Comprehensive testing
- [ ] Documentation updates
- [ ] Build scripts
- [ ] Installation instructions
- [ ] Usage examples

## Error Handling Strategy
- Validate all user inputs (negative amounts, empty descriptions)
- Handle file I/O errors gracefully
- Provide clear, actionable error messages
- Implement proper exit codes
- Handle edge cases (non-existent IDs, invalid dates)

## Data Storage Format
```json
{
  "expenses": [
    {
      "id": 1,
      "date": "2024-08-06T00:00:00Z",
      "description": "Lunch",
      "amount": 20.00,
      "category": "Food"
    }
  ],
  "nextId": 2
}
```

## Testing Strategy
- Unit tests for all core functions
- Integration tests for CLI commands
- Test data validation
- Test file operations
- Test edge cases and error conditions

## Success Criteria
- All required commands work as specified
- Data persists between sessions
- Proper error handling for invalid inputs
- Clean, maintainable code structure
- Comprehensive test coverage
- Clear documentation and usage instructions
