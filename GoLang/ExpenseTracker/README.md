# Expense Tracker

A simple command-line expense tracker application built in Go to manage your finances. The application allows users to add, delete, and view their expenses, and provides summaries of expenses.

## Features

### Core Features

- ✅ Add an expense with a description and amount
- ✅ Update an expense
- ✅ Delete an expense
- ✅ View all expenses
- ✅ View a summary of all expenses
- ✅ View a summary of expenses for a specific month (of current year)

### Additional Features (Future Enhancements)

- Add expense categories and allow users to filter expenses by category
- Allow users to set a budget for each month and show a warning when the user exceeds the budget
- Allow users to export expenses to a CSV file

## Usage

### Commands

```bash
# Add an expense
$ expense-tracker add --description "Lunch" --amount 20
# Expense added successfully (ID: 1)

$ expense-tracker add --description "Dinner" --amount 10
# Expense added successfully (ID: 2)

# List all expenses
$ expense-tracker list
# ID  Date       Description  Amount
# 1   2024-08-06  Lunch        $20
# 2   2024-08-06  Dinner       $10

# View summary
$ expense-tracker summary
# Total expenses: $30

# Delete an expense
$ expense-tracker delete --id 2
# Expense deleted successfully

# View summary again
$ expense-tracker summary
# Total expenses: $20

# View monthly summary
$ expense-tracker summary --month 8
# Total expenses for August: $20
```

## Implementation Details

- Built with Go
- Uses JSON file storage (`expenses.json`)
- Command-line argument parsing using Go's `flag` package
- Error handling for invalid inputs and edge cases (e.g., negative amounts, non-existent expense IDs)
- Modular code structure with separate files for commands, storage, and data models

## Project Structure

```
ExpenseTracker/
├── README.md          # This file
├── go.mod             # Go module file
├── main.go            # Entry point and routing
├── commands.go        # Command handlers
├── expense.go         # Expense data structure
└── storage.go         # File I/O operations
```

## Getting Started

1. Clone the repository
2. Build the application:
   ```bash
   go build
   ```
3. Run the application:
   ```bash
   go run .
   ```

This project is a great way to practice logic building skills and learn how to interact with the filesystem using a CLI application. It also helps understand how to manage data and provide useful information to users in a structured way.

