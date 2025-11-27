const dataFile = "expenses.json"

import (
    "encoding/json"
    "os"
)

func loadExpenses() []Expense {
    data, err := os.ReadFile(dataFile)
	// If file is missing or there is an error, return empty slice
    if err != nil {
        return []Expense{}  
    }
    
    var expenses []Expense
    if err := json.Unmarshal(data, &expenses); err != nil {
        return []Expense{}  // If JSON is corrupted, return empty slice
    }
    
    return expenses
}

func saveExpenses(expenses []Expense) error {
    data, err := json.MarshalIndent(expenses, "", "  ")  // formatted JSON
    if err != nil {
        return err
    }
    return os.WriteFile(dataFile, data, 0644)
}