package models

import "github.com/uadmin/uadmin"

// TaxType !
type TaxType int

// Exclusive !
func (TaxType) Exclusive() TaxType {
	return 1
}

// Inclusive !
func (TaxType) Inclusive() TaxType {
	return 2
}

type Tax struct {
	uadmin.Model
	Name string  `uadmin:"multilingual;required;search"`
	Type TaxType `uadmin:"help:Inclusive means the tax is included already in the price"`
	Rate float64 `uadmin:"min:0;help:For 12% use 0.12"`
}

func (t *Tax) String() string {
	return t.Name
}
