package main

import "strconv"

type Data struct {
	incomes  []float64
	expenses []float64
}

func NewData() *Data {
	return &Data{}
}

func (d *Data) AddIncome(amount string) {
	if value, err := strconv.ParseFloat(amount, 64); err == nil {
		d.incomes = append(d.incomes, value)
	}
}

func (d *Data) AddExpense(amount string) {
	if value, err := strconv.ParseFloat(amount, 64); err == nil {
		d.expenses = append(d.expenses, value)
	}
}

func (d *Data) CalculateBalance() string {
	var totalIncome, totalExpense float64
	for _, income := range d.incomes {
		totalIncome += income
	}
	for _, expense := range d.expenses {
		totalExpense += expense
	}
	balance := totalIncome - totalExpense
	return strconv.FormatFloat(balance, 'f', 2, 64)
}
