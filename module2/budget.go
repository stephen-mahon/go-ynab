package budget

import (
	"errors"
	"time"
)

type Budget struct {
	Max   float32
	Items []Item
}

type Item struct {
	Description string
	Price       float32
}

var report map[time.Month]*Budget

func InitializeReport() {
	report = make(map[time.Month]*Budget)
}

func init() {
	InitializeReport()
}

func (b Budget) CurrentCost() float32 {
	var sum float32
	for _, item := range b.Items {
		sum += item.Price
	}
	return sum
}

var errDoesNotFitBudget = errors.New("item does not fit the budget")

var errReportIsFull = errors.New("report is full")

var errDuplicateEntry = errors.New("cannot add duplicate entry")

func (b *Budget) AddItem(description string, price float32) error {
	if b.CurrentCost()+price > b.Max {
		return errDoesNotFitBudget
	}
	newItem := Item{Description: description, Price: price}
	b.Items = append(b.Items, newItem)
	return nil
}

func (b *Budget) RemoveItem(description string) {
	for i := range b.Items {
		if b.Items[i].Description == description {
			b.Items = append(b.Items[:i], b.Items[i+1:]...)
			break
		}
	}
}

func CreateBudget(month time.Month, max float32) (*Budget, error) {
	if len(report) >= 12 {
		return nil, errReportIsFull
	}

	_, hasEntry := report[month]

	if hasEntry {
		return nil, errDuplicateEntry
	}

	var newBudget = &Budget{Max: max}
	report[month] = newBudget
	return newBudget, nil
}
