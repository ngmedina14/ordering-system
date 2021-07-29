package models

import (
	"strings"
	"time"

	"github.com/uadmin/uadmin"
)

// ItemStatus !
type ItemStatus int

// Preparing !
func (ItemStatus) Pending() ItemStatus {
	return 1
}

// Ready !
func (ItemStatus) Ready() ItemStatus {
	return 2
}

// Served !
func (ItemStatus) Served() ItemStatus {
	return 3
}

// Cancelled !
func (ItemStatus) Cancelled() ItemStatus {
	return 4
}

type OrderItem struct {
	uadmin.Model
	Order          Order `uadmin:"list_exclude"`
	OrderID        uint
	MenuItem       MenuItem `uadmin:"list_exclude"`
	MenuItemID     uint
	ItemOption     []ItemOption `uadmin:"list_exclude"`
	ItemOptionList string       `uadmin:"read_only"`
	Quantity       int          `uadmin:"min:0"`
	Status         ItemStatus   `uadmin:"help:Status are Preparing, Ready, Served, and Cancelled"`
	//TaxID
	ItemGrossSales float64    `uadmin:"min:0;help:Gross sales are the value of all of Menu Item sumation without accounting for any deductions."`
	ItemNetTotal   float64    `uadmin:"read_only;search;min:0;help:Net Total are gross sales minus deductions: discounts."`
	ServedTime     *time.Time `uadmin:"filter"`
}

//Save the Names of all Selected Item Option and put it in ItemOptionList
func (i *OrderItem) Save() {
	// Add a new string array type variable called categoryList
	itemList := []string{}

	// Append every element to the categoryList array
	for c := range i.ItemOption {
		itemList = append(itemList, i.ItemOption[c].Name)
	}

	// Concatenate the categoryList to a single string separated by comma
	joinList := strings.Join(itemList, ", ")

	// Store the joined string to the CategoryList field
	i.ItemOptionList = joinList

	// Save it to the database
	uadmin.Save(i)
}
