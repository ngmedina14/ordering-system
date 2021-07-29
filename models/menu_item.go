// Menu Item Table
// *Name
// *Description
// *Image
// *CategoryID
// *Price
// *Dine In
// *Take Out
// *Availability
// *Hide Everywhere
// *Hide on Web
// *Featured
// *Code
// *DiscountID //Preset purpose
// *Time Availability
// *Position
// *TaxID
// *ModifierID
// *ItemOptionID
package models

import "github.com/uadmin/uadmin"

//MenuItem is the product of the restaurant -- and After Category in the web
type MenuItem struct {
	uadmin.Model
	Name             string
	Description      string
	Image            string       `uadmin:"image"`
	MenuCategory     MenuCategory `uadmin:"list_exclude;help:Select the appropriate Category in this Menu Item."`
	MenuCategoryID   uint
	Price            float64 `uadmin:"min:0"`
	ShowDineIn       bool
	ShowTakeOut      bool
	TimeAvailability string `uadmin:"help:Empty for all day availability. For 7:30AM-1PM and 4PM-9PM use:7:30AM-1PM,4PM-9PM.;"`
	HideEverywhere   bool   `uadmin:"filter;help:Hide on cashier and web."`
	HideOnWeb        bool   `uadmin:"filter;help:Hide on web only."`
	Featured         bool   `uadmin:"filter;help:Show at the slide."`
	Code             string `uadmin:"list_exclude;help:Special word that is used for this item. For Main Course, use MC1 on the first item, MC2 on the second item. For Beverage, use BV1 on the first item, BV2 on the second item. If you are not sure, leave it blank."`
	// *DiscountID //Preset purpose
	Position int `uadmin:"list_exclude;help:To show this item at the top of the list, use '1'. To show this item in the third position from the top, use '3'."`
	// *TaxID
	OptionCategory   OptionCategory `uadmin:"list_exclude;help:Select the appropriate Item Option in the choices for this Option Category."`
	OptionCategoryID uint
}

func (mi MenuItem) String() string {
	return mi.Name
}
