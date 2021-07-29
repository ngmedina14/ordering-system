package models

import (
	"time"

	"github.com/uadmin/uadmin"
)

type OfficialReceipt struct {
	uadmin.Model
	CreatedAt time.Time
	// Order Order
	// OrderID uint
	// Cashier Employee
	// CashierID uint
	Voided   bool
	VoidDate *time.Time
}

func (o *OfficialReceipt) Save() {
	uadmin.Save(o)
}

func (o *OfficialReceipt) Void() {
	o.Voided = true
	now := time.Now()
	o.VoidDate = &now
	uadmin.Save(o)
}
