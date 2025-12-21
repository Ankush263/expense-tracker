package controllers

import (
	"encoding/json"
	"os"

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
