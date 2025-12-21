package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ankush263/expense-tracker/internal/controllers"
	"github.com/ankush263/expense-tracker/utils"
)

// PLANNING....

// $ expense-tracker add --description "Lunch" --amount 20
// # Expense added successfully (ID: 1)

// $ expense-tracker add --description "Dinner" --amount 10
// # Expense added successfully (ID: 2)

// $ expense-tracker list
// # ID  Date       Description  Amount
// # 1   2024-08-06  Lunch        $20
// # 2   2024-08-06  Dinner       $10

// $ expense-tracker summary
// # Total expenses: $30

// $ expense-tracker delete --id 2
// # Expense deleted successfully

// $ expense-tracker summary
// # Total expenses: $20

// $ expense-tracker summary --month 8
// # Total expenses for August: $20


func main() {
	args := os.Args[1:]
	
	if len(args) == 0 {
		fmt.Println("No commands provided")
		return
	}

	switch(args[0]) {
	case "list":
		if len(args) > 1 {
			fmt.Println("Usage: expense-tracket-cli list")
			return
		}
		data := controllers.GetAllExpenses()
		structureData, err := json.MarshalIndent(data, "", " ")
		utils.CheckNilError(err)
		fmt.Println(string(structureData))
		return
	}
	
}