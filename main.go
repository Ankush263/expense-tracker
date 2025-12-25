package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

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

	case "add":
		if len(args) < 2 || len(args) > 5 {
			fmt.Println("Usage: expense-tracket-cli add --description 'Lunch' --amount 20")
			return
		}

		var description string
		var amount int
		
		switch(args[1]) {
		case "--description":
			description = args[2]

		case "--amount":
			intamount, err := strconv.Atoi(args[2])
			amount = intamount
			utils.CheckNilError(err)
		}

		switch(args[3]) {
		case "--description":
			description = args[4]

		case "--amount":
			intamount, err := strconv.Atoi(args[4])
			amount = intamount
			utils.CheckNilError(err)
		}

		newExpense := controllers.AddToExpanses(description, amount)

		data, err := json.MarshalIndent(newExpense, "", " ")
		utils.CheckNilError(err)

		fmt.Println(string(data))
		return 

	case "summary":
		if len(args) > 1 {
			fmt.Println("Usage: expense-tracket-cli summary")
			return
		}

		totalExpanse := controllers.GetTotalExpanse()

		fmt.Println("$", totalExpanse)
		return
	}
}