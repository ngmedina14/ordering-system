package models

import "github.com/uadmin/uadmin"

//Menu Category is a Category when selecting a Menu Item on web or cashier
type MenuCategory struct {
	uadmin.Model
	Name        string
	Description string
	Image       string        `uadmin:"image"`
	Parent      *MenuCategory `uadmin:"display_name:Main Category;help:Select the existing category as a parent in the choices for this category.;filter"`
	ParentID    uint
	HiddenOnWeb bool
}

func (mc *MenuCategory) String() string {
	return mc.Name
}
