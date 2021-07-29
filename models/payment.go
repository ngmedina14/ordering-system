package models

import (
	"time"

	"github.com/uadmin/uadmin"
)

type PaymentType int

func (p PaymentType) Cash() PaymentType {
	return 1
}

func (p PaymentType) Card() PaymentType {
	return 2
}

func (p PaymentType) Other() PaymentType {
	return 3
}

func (p PaymentType) Gcash() PaymentType {
	return 4
}

func (p PaymentType) Paymaya() PaymentType {
	return 5
}

func (p PaymentType) Paypal() PaymentType {
	return 5
}

type Payment struct {
	uadmin.Model
	Order       Order
	OrderID     uint
	CreatedAt   time.Time `uadmin:"read_only"`
	Type        PaymentType
	Amount      float64  `uadmin:"help:Amount of the payment of the Customer"`
	Change      float64  `uadmin:"help:change is amount minus Order Charge"`
	Cashier     Employee `uadmin:"list_exclude"`
	CashierID   uint
	Transferred bool
}
