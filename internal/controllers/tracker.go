package controllers

import (
	"encoding/json"
	"errors"
	"os"
	"time"

	"github.com/ankush263/expense-tracker/internal/model"
	"github.com/ankush263/expense-tracker/utils"
)

const filepath string = "data/expense-tracker.json"

func LoadExpenses() []model.ExpenseTracker {
	// Check if the file does not exists in filepath then return empty task
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return []model.ExpenseTracker{}
	}

	data, err := os.ReadFile(filepath)
	utils.CheckNilError(err)

	if len(data) == 0 {
		return []model.ExpenseTracker{}
	}

	var trackerData []model.ExpenseTracker
	err = json.Unmarshal(data, &trackerData)
	utils.CheckNilError(err)

	return trackerData
}

func GetAllExpenses() []model.ExpenseTracker {
	expenses := LoadExpenses()

	return expenses
}

func SaveExpenses(expenses []model.ExpenseTracker) error {
	data, err := json.MarshalIndent(expenses, "", " ")
	utils.CheckNilError(err)

	return os.WriteFile(filepath, data, 0644)
}

func AddToExpanses(description string, amount int) model.ExpenseTracker {
	expenses := LoadExpenses()

	var newId int = 1
	if len(expenses) > 0 {
		newId = expenses[len(expenses) - 1].ID + 1
	}

	newExpense := model.ExpenseTracker{
		ID: newId,
		CreatedAt: time.Now(),
		Description: description,
		Amount: amount,
	}

	expenses = append(expenses, newExpense)

	err := SaveExpenses(expenses)
	utils.CheckNilError(err)

	return  newExpense
}

func GetExpenseById(id int) model.ExpenseTracker {
	expense := LoadExpenses()

	if len(expense) == 0 {
		return model.ExpenseTracker{}
	}

	for _, exp := range(expense) {
		if exp.ID == id {
			return exp
		}
	}

	return model.ExpenseTracker{}
}

func GetTotalExpanse(month *int) int {
	expenses := LoadExpenses()

	if len(expenses) == 0 {
		return 0
	}

	var total int = 0

	if (month == nil) {
		for _, expense := range(expenses) {
			total += expense.Amount
		}
	} else {
		for _, expense := range(expenses) {
			if int(expense.CreatedAt.Month()) == *month {
				total += expense.Amount
			}
		}
	}

	return total
}

func UpdateExpanseDescription(id int, description string) (model.ExpenseTracker, error) {
	expenses := LoadExpenses()

	if len(expenses) == 0 {
		return model.ExpenseTracker{}, errors.New("No expenses found")
	}

	for i, expense := range(expenses) {
		if expense.ID == id {
			expenses[i].Description = description

			err := SaveExpenses(expenses)
			utils.CheckNilError(err)

			return expenses[i], nil
		}
	}

	return model.ExpenseTracker{}, errors.New("Failed to update expense")
}

func UpdateExpense(id int, description *string, amount *int) (model.ExpenseTracker, error) {
	expenses := LoadExpenses()

	if len(expenses) == 0 {
		return model.ExpenseTracker{}, errors.New("No expenses found")
	}

	for i, expense := range(expenses) {
		if expense.ID == id {
			if description != nil {
				expenses[i].Description = *description
			}

			if amount != nil {
				expenses[i].Amount = *amount
			}

			err := SaveExpenses(expenses)
			utils.CheckNilError(err)

			return expenses[i], nil
		}
	}

	return model.ExpenseTracker{}, errors.New("Failed to update expense amount")
}

func DeleteExpense(id int) error {
	expenses := LoadExpenses()

	if len(expenses) == 0 {
		return errors.New("Expenses are empty")
	}

	for i, expense := range(expenses) {
		if expense.ID == id {
			expenses = append(expenses[:i], expenses[i+1:]...)

			err := SaveExpenses(expenses)
			utils.CheckNilError(err)
			return nil
		}

	}

	return errors.New("Expense deletion failed")
}
