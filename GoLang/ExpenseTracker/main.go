package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Args[0] is the program name
	// os.Args[1] is the first argument (the command)
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	// Get the command from command-line arguments
	command := os.Args[1]

	// Handle different commands
	switch command {
	case "add":
		handleAdd()
	case "list":
		handleList()
	case "delete":
		handleDelete()
	case "update":
		handleUpdate()
	case "summary":
		handleSummary()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
	}
}

// printUsage displays help information to the user
func printUsage() {
	fmt.Println(`Usage: expense-tracker <command> [options]

Commands:
  add       - Add a new expense
  list      - List all expenses
  delete    - Delete an expense
  update    - Update an expense
  summary   - Show expense summary`)
}

