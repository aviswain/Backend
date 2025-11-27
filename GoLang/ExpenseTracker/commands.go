package main

import "fmt"

// handleAdd adds a new expense
func handleAdd() {
	fmt.Println("Add command - to be implemented")
	// TODO: Parse --description and --amount flags
	// TODO: Create expense and save it
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

