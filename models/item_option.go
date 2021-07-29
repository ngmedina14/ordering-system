package models

import "github.com/uadmin/uadmin"

//Item Option is after selecting an option category -- example if the option category is 'Size' then the Item option is Large, Medium, Small
type ItemOption struct {
	uadmin.Model
	OptionCategory   OptionCategory `uadmin:"list_exclude;help:Select the appropriate Item Option in the choices for this Option Category."`
	OptionCategoryID uint
	Name             string
	Price            float64 `uadmin:"min:0"`
}

func (io ItemOption) String() string {
	return io.Name
}
