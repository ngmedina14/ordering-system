package models

import "github.com/uadmin/uadmin"

//Discount will be applied in Order
type Discount struct {
	uadmin.Model
	Name           string `uadmin:"multilingual;required;search"`
	Percentage     bool
	ApplyBeforeTax bool
	Value          float64 `uadmin:"min:0"`
	// ExcludeTax     []Tax
}

func (d *Discount) String() string {
	return d.Name
}
