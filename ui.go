package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type UI struct {
	window fyne.Window
	data   *Data
}

func NewUI(window fyne.Window) *UI {
	return &UI{window: window, data: NewData()}
}

func (ui *UI) MakeUI() fyne.CanvasObject {
	incomeLabel := widget.NewLabel("Income")
	incomeEntry := widget.NewEntry()
	addIncomeButton := widget.NewButton("Add Income", func() {
		ui.data.AddIncome(incomeEntry.Text)
		incomeEntry.SetText("")
	})

	expenseLabel := widget.NewLabel("Expense")
	expenseEntry := widget.NewEntry()
	addExpenseButton := widget.NewButton("Add Expense", func() {
		ui.data.AddExpense(expenseEntry.Text)
		expenseEntry.SetText("")
	})

	balanceLabel := widget.NewLabel("Balance: $0")
	updateBalanceButton := widget.NewButton("Update Balance", func() {
		balance := ui.data.CalculateBalance()
		balanceLabel.SetText("Balance: $" + balance)
	})

	content := container.NewVBox(
		incomeLabel,
		incomeEntry,
		addIncomeButton,
		expenseLabel,
		expenseEntry,
		addExpenseButton,
		balanceLabel,
		updateBalanceButton,
	)

	return content
}
