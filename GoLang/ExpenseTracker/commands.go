package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

// handleAdd adds a new expense
func handleAdd() {
	// Create a new flag set to parse flags for the "add" command
	addFlags := flag.NewFlagSet("add", flag.ExitOnError)

	// Define flags
	description := addFlags.String("description", "", "Expense description (required)")
	amount := addFlags.Float64("amount", 0, "Expense amount (required)")

	// Parse flags from os.Args[2:] (skip program name and "add" command)
	if err := addFlags.Parse(os.Args[2:]); err != nil {
		fmt.Printf("Error parsing flags: %v\n", err)
		return
	}

	if *description == "" {
		fmt.Println("Error: --description is required")
		addFlags.Usage()
		return
	}
	if *amount <= 0 {
		fmt.Println("Error: --amount must be positive")
		addFlags.Usage()
		return
	}

	date := time.Now().Format("2006-01-02")

	expenses := loadExpenses()

	//  next ID  = max ID + 1
	nextID := 1
	if len(expenses) > 0 {
		maxID := 0
		for _, exp := range expenses {
			if exp.ID > maxID {
				maxID = exp.ID
			}
		}
		nextID = maxID + 1
	}

	newExpense := Expense{
		ID:          nextID,
		Date:        date,
		Description: *description,
		Amount:      *amount,
	}


	expenses = append(expenses, newExpense)

	if err := saveExpenses(expenses); err != nil {
		fmt.Printf("Error saving expense: %v\n", err)
		return
	}

	fmt.Printf("Expense added successfully (ID: %d)\n", nextID)
}

// handleList displays all expenses
func handleList() {
	fmt.Println("List command - to be implemented")
	// TODO: Load expenses from file
	// TODO: Display all expenses in a table format
}

// handleDelete removes an expense by ID
func handleDelete() {
	fmt.Println("Delete command - to be implemented")
	// TODO: Parse --id flag
	// TODO: Delete expense by ID
}

// handleUpdate modifies an existing expense
func handleUpdate() {
	fmt.Println("Update command - to be implemented")
	// TODO: Parse --id, --description, and --amount flags
	// TODO: Update existing expense
}

// handleSummary shows expense totals
func handleSummary() {
	fmt.Println("Summary command - to be implemented")
	// TODO: Load expenses and calculate total
	// TODO: Support --month flag for monthly summary
}

