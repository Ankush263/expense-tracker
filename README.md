# 💸 Expense Tracker CLI

A simple, fast, and minimal **command-line expense tracker** to manage daily expenses directly from your terminal.

This tool helps you record, update, delete, and summarize expenses without any external dependencies or databases. All data is stored locally, keeping things lightweight and transparent.

---

## ✨ What This Tool Does

- Track expenses with description and amount
- Update or delete existing records
- View all expenses in a clean list
- Get total expense summaries
- View monthly expense breakdowns
- Works completely offline
- Stores data locally in a file

Perfect for developers who prefer **terminal tools over apps** and want a quick way to track spending.

---

## 📦 Features

### ✅ Expense Management
- Add a new expense
- Update an existing expense
- Delete an expense by ID

### 📋 Viewing & Insights
- List all recorded expenses
- View total expenses
- View expenses for a specific month

### 🗂 Data Handling
- Local file-based storage
- No database required
- Simple and readable data format

---

## 🧪 Example Usage

```bash
$ expense-tracker add --description "Coffee" --amount 5
Expense added successfully (ID: 1)

$ expense-tracker add --description "Groceries" --amount 30
Expense added successfully (ID: 2)

$ expense-tracker list
ID  Date        Description  Amount
1   2024-08-06  Coffee       $5
2   2024-08-06  Groceries    $30

$ expense-tracker summary
Total expenses: $35

$ expense-tracker summary --month 8
Total expenses for August: $35

$ expense-tracker delete --id 1
Expense deleted successfully
