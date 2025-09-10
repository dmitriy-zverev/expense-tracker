# 💰 Expense Tracker CLI

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-blue?style=for-the-badge)](LICENSE)
[![Build Status](https://img.shields.io/badge/Build-Passing-brightgreen?style=for-the-badge)](https://github.com/dmitriy-zverev/expense-tracker)
[![Coverage](https://img.shields.io/badge/Coverage-85%25-green?style=for-the-badge)](https://github.com/dmitriy-zverev/expense-tracker)

> A powerful, lightweight command-line expense tracking application built with Go. Manage your personal finances with ease through an intuitive CLI interface.

## ✨ Features

- 🚀 **Lightning Fast** - Built with Go for optimal performance
- 💾 **Local Storage** - Your data stays on your machine (JSON-based)
- 📊 **Budget Management** - Set and track monthly budgets by category
- 📈 **Smart Summaries** - Detailed expense analytics and reporting
- 📤 **CSV Export** - Export your data for external analysis
- 🔍 **Advanced Filtering** - Filter by category, month, or date range
- ✅ **Comprehensive Testing** - 85%+ test coverage for reliability
- 🛡️ **Input Validation** - Robust error handling and data validation

## 🚀 Quick Start

### Installation

```bash
# Clone the repository
git clone https://github.com/dmitriy-zverev/expense-tracker.git
cd expense-tracker

# Build the application
go build -o expense-tracker

# Make it executable (Unix/Linux/macOS)
chmod +x expense-tracker

# Optional: Add to PATH for global access
sudo mv expense-tracker /usr/local/bin/
```

### Basic Usage

```bash
# Add your first expense
expense-tracker add --amount 25.50 --description "Coffee and pastry" --category "Food"

# List all expenses
expense-tracker list

# View monthly summary
expense-tracker summary --month 9

# Set a budget
expense-tracker budget --month 9 --category "Food" --limit 500

# Export to CSV
expense-tracker export
```

## 📖 Documentation

### Core Commands

#### 💸 Adding Expenses

```bash
# Basic expense
expense-tracker add --amount 15.99 --description "Lunch"

# With category
expense-tracker add --amount 50.00 --description "Gas" --category "Transportation"

# Multiple expenses quickly
expense-tracker add -a 12.50 -d "Coffee" -c "Food"
expense-tracker add -a 8.99 -d "Parking" -c "Transportation"
```

#### 📋 Listing Expenses

```bash
# List all expenses
expense-tracker list

# Filter by category
expense-tracker list --category "Food"

# Filter by month
expense-tracker list --month 9

# Combine filters
expense-tracker list --category "Food" --month 9
```

#### ✏️ Managing Expenses

```bash
# Update an expense
expense-tracker update --id 1 --amount 18.99 --description "Updated lunch"

# Delete an expense
expense-tracker delete --id 1
```

#### 📊 Analytics & Summaries

```bash
# Overall summary
expense-tracker summary

# Monthly summary
expense-tracker summary --month 9

# Category-specific summary
expense-tracker summary --category "Food"

# Combined filters
expense-tracker summary --month 9 --category "Food"
```

#### 💰 Budget Management

```bash
# Set a monthly budget
expense-tracker budget --month 9 --category "Food" --limit 500.00

# List all budgets
expense-tracker budget --list

# Remove a budget
expense-tracker budget --month 9 --category "Food" --remove
```

#### 📤 Data Export

```bash
# Export all data to CSV
expense-tracker export

# Export specific month
expense-tracker export --month 9

# Custom output file
expense-tracker export --output my-expenses.csv
```

### Command Reference

| Command | Description | Options |
|---------|-------------|---------|
| `add` | Add a new expense | `--amount`, `--description`, `--category` |
| `list` | List expenses | `--category`, `--month` |
| `update` | Update existing expense | `--id`, `--amount`, `--description`, `--category` |
| `delete` | Delete an expense | `--id` |
| `summary` | Show expense summary | `--month`, `--category` |
| `budget` | Manage budgets | `--month`, `--category`, `--limit`, `--list`, `--remove` |
| `export` | Export to CSV | `--month`, `--output` |
| `help` | Show help information | - |

## 🏗️ Architecture

### Project Structure

```
expense-tracker/
├── 📁 cmd/                    # CLI command implementations
│   ├── add.go                 # Add expense command
│   ├── budget.go              # Budget management commands
│   ├── delete.go              # Delete expense command
│   ├── export.go              # CSV export functionality
│   ├── list.go                # List expenses command
│   ├── root.go                # Root command and CLI setup
│   ├── summary.go             # Summary and analytics
│   └── update.go              # Update expense command
├── 📁 internal/               # Internal application logic
│   ├── 📁 budget/             # Budget management
│   │   ├── budget.go          # Budget operations
│   │   └── budget_test.go     # Budget tests
│   ├── 📁 expense/            # Expense management
│   │   ├── expense.go         # Core expense operations
│   │   └── expense_test.go    # Expense tests
│   ├── 📁 storage/            # Data persistence layer
│   │   ├── file.go            # File-based storage
│   │   └── storage_test.go    # Storage tests
│   └── 📁 utils/              # Utility functions
│       ├── validation.go      # Input validation
│       └── validation_test.go # Validation tests
├── 📁 data/                   # Data storage
│   ├── expenses.json          # Expense data
│   └── budgets.json           # Budget configuration
├── main.go                    # Application entry point
├── go.mod                     # Go module definition
└── README.md                  # This file
```

### Data Models

#### Expense Model
```go
type Expense struct {
    ID          int       `json:"id"`
    Amount      float64   `json:"amount"`
    Date        time.Time `json:"date"`
    Description string    `json:"description"`
    Category    string    `json:"category"`
    Month       int       `json:"month"`
    IsDeleted   bool      `json:"is_deleted"`
}
```

#### Budget Model
```go
type Budget struct {
    Month    int     `json:"month"`
    Year     int     `json:"year"`
    Category string  `json:"category"`
    Limit    float64 `json:"limit"`
}
```

### Design Principles

- **🎯 Single Responsibility** - Each package has a clear, focused purpose
- **🔒 Encapsulation** - Internal packages protect implementation details
- **🧪 Testability** - Comprehensive test suite with high coverage
- **⚡ Performance** - Efficient file I/O and minimal memory footprint
- **🛡️ Reliability** - Robust error handling and input validation

## 🧪 Testing

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test ./... -cover

# Run tests with verbose output
go test ./... -v

# Run specific package tests
go test ./internal/expense -v
```

### Test Coverage

| Package | Coverage | Description |
|---------|----------|-------------|
| `internal/utils` | 100% | Input validation and utilities |
| `internal/storage` | 89% | File operations and data persistence |
| `internal/budget` | 85% | Budget management functionality |
| `internal/expense` | 82% | Core expense operations |

### Test Categories

- **Unit Tests** - Individual function testing
- **Integration Tests** - Multi-component workflows
- **Edge Case Tests** - Boundary conditions and error scenarios
- **Performance Tests** - Benchmarks for critical operations

## 🔧 Development

### Prerequisites

- Go 1.21 or higher
- Git

### Development Setup

```bash
# Clone the repository
git clone https://github.com/dmitriy-zverev/expense-tracker.git
cd expense-tracker

# Install dependencies
go mod download

# Run tests
go test ./...

# Build for development
go build -o expense-tracker-dev

# Run with sample data
./expense-tracker-dev add --amount 10.99 --description "Test expense"
```

### Building for Production

```bash
# Build optimized binary
go build -ldflags="-s -w" -o expense-tracker

# Cross-compile for different platforms
GOOS=linux GOARCH=amd64 go build -o expense-tracker-linux
GOOS=windows GOARCH=amd64 go build -o expense-tracker.exe
GOOS=darwin GOARCH=amd64 go build -o expense-tracker-macos
```

### Code Quality

```bash
# Format code
go fmt ./...

# Lint code (requires golangci-lint)
golangci-lint run

# Vet code
go vet ./...

# Security scan (requires gosec)
gosec ./...
```

## 📊 Examples

### Daily Expense Tracking

```bash
# Morning coffee
expense-tracker add -a 4.50 -d "Morning coffee" -c "Food"

# Lunch
expense-tracker add -a 12.99 -d "Sandwich and drink" -c "Food"

# Transportation
expense-tracker add -a 2.75 -d "Bus fare" -c "Transportation"

# Check daily total
expense-tracker summary
```

### Monthly Budget Management

```bash
# Set monthly budgets
expense-tracker budget --month 9 --category "Food" --limit 400
expense-tracker budget --month 9 --category "Transportation" --limit 150
expense-tracker budget --month 9 --category "Entertainment" --limit 200

# Check budget status
expense-tracker budget --list

# View spending vs budget
expense-tracker summary --month 9
```

### Data Analysis

```bash
# Export for spreadsheet analysis
expense-tracker export --output september-expenses.csv

# Category breakdown
expense-tracker summary --category "Food"
expense-tracker summary --category "Transportation"
expense-tracker summary --category "Entertainment"

# Monthly comparison
expense-tracker summary --month 8
expense-tracker summary --month 9
```

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guidelines](CONTRIBUTING.md) for details.

### Quick Contribution Guide

1. **Fork** the repository
2. **Create** a feature branch (`git checkout -b feature/amazing-feature`)
3. **Commit** your changes (`git commit -m 'Add amazing feature'`)
4. **Push** to the branch (`git push origin feature/amazing-feature`)
5. **Open** a Pull Request

### Development Guidelines

- Write tests for new functionality
- Follow Go conventions and best practices
- Update documentation for user-facing changes
- Ensure all tests pass before submitting PR

## 📄 License

This project is free to use. No additional requirements.

## 🙏 Acknowledgments

- Built with ❤️ using [Go](https://golang.org/)
- Inspired by the need for simple, effective personal finance management
- Thanks to the Go community for excellent tooling and libraries

---

<div align="center">

**[⬆ Back to Top](#-expense-tracker-cli)**

Made with ❤️ by [Dmitriy Zverev](https://github.com/dmitriy-zverev)

</div>
