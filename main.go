package main

import (
	"fmt"
	budget "go-ynab/module2"
	"time"
)

var months = []time.Month{
	time.January,
	time.February,
	time.March,
	time.April,
	time.May,
	time.June,
	time.July,
	time.August,
	time.September,
	time.October,
	time.November,
	time.December,
}

func main() {
	bu, _ := budget.CreateBudget(time.January, 1000)
	bu.AddItem("bananas", 10.0)

	fmt.Println("Items in January:", len(bu.Items))

	budget.CreateBudget(time.February, 1000)

	bu2 := budget.GetBudget(time.February)
	bu2.AddItem("bananas", 10.0)
	bu2.AddItem("coffee", 3.99)
	bu2.AddItem("gym", 50.0)
	bu2.RemoveItem("coffee")
	fmt.Println("Items in February:", len(bu2.Items))
	fmt.Printf("Current cost for February: $%.2f \n", bu2.CurrentCost())

	fmt.Println("Resetting Budget Report...")
	budget.InitializeReport()

	for _, month := range months {
		_, err := budget.CreateBudget(month, 100.00)
		if err == nil {
			fmt.Println("Budget created for", month)
		} else {
			fmt.Println("Error creating budget:", err)
		}
	}

	_, err := budget.CreateBudget(time.December, 100.00)
	fmt.Println(err)

}
