package models

import (
	"time"

	"github.com/uadmin/uadmin"
)

type OrderType int

func (OrderType) DineIn() OrderType {
	return 1
}

func (OrderType) TakeOut() OrderType {
	return 2
}

func (OrderType) Delivery() OrderType {
	return 3
}

func (OrderType) PickUp() OrderType {
	return 4
}

// OrderStatus !
type OrderStatus int

// Pending !
func (OrderStatus) Pending() OrderStatus {
	return 1
}

// Ready !
func (OrderStatus) Ready() OrderStatus {
	return 2
}

// Served !
func (OrderStatus) Served() OrderStatus {
	return 3
}

// Cancelled !
func (OrderStatus) Cancelled() OrderStatus {
	return 4
}

//Order is the transaction made by the customer
type Order struct {
	uadmin.Model
	Customer        Customer `uadmin:"list_exclude;help:Select the Customer. If the customer has different mac address create new customer.;filter"`
	CustomerID      uint
	CreatedAt       *time.Time `uadmin:"filter"`
	OrderType       OrderType
	DeliveryCharge  float64     `uadmin:"list_exclude;min:0"`
	GrossSales      float64     `uadmin:"min:0;help:Gross sales are the value of all of Menu Item sumation without accounting for any deductions."`
	NetTotal        float64     `uadmin:"read_only;search;min:0;help:Net Total are gross sales minus deductions: discounts."`
	Tax             float64     `uadmin:"read_only;list_exclude;min:0;help:Show amount of the tax included or excluded"`
	Charge          float64     `uadmin:"read_only;list_exclude;min:0;help:Amount that customer need to pay-- sum of Net Total, Tax, Delivery Charge"`
	Change          float64     `uadmin:"list_exclude;min:0;help:Amount Paid minus charge"`
	Paid            bool        `uadmin:"filter;help:Check if the customer has paid his order"`
	Status          OrderStatus `uadmin:"filter;help:Status are Pending, Ready, Served, and Cancelled"`
	Waiter          Employee    `uadmin:"list_exclude"`
	WaiterID        uint
	IsEcash         bool `uadmin:"hidden;list_exclude"`
	ReferenceNumber string
	// *PrintCount
	// *Reprint
}
