package models

import "github.com/uadmin/uadmin"

//Option Category is Category before choosing an Item Option -- Example: Size, Variant, Level, adds on
type OptionCategory struct {
	uadmin.Model
	Name        string `uadmin:"list_exclude;help:While creating an item this name will be displayed."`
	DisplayName string `uadmin:"help:Will be Display on Website as Option Category for the specific Menu Item"`
	ChoiceType  bool   // One or Multiple
}

func (oc OptionCategory) String() string {
	return oc.Name
}
