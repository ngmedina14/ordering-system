package models

import "github.com/uadmin/uadmin"

//Combo Item is a combination of Menu Item in a single product and priced as one hence it will be treated as a Menu Item
type ComboItem struct {
	uadmin.Model
	Name         string
	Description  string
	Price        float64    `uadmin:"min:0"`
	Image        string     `uadmin:"image"`
	MenuItem     []MenuItem `uadmin:"list_exclude"`
	MenuItemList string     `uadmin:"read_only"`
}

func (ci ComboItem) String() string {
	return ci.Name
}
