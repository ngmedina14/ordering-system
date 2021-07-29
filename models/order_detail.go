package models

import (
	"time"

	"github.com/uadmin/uadmin"
)

//Todo: Identify if ill put Gross sale and Net Total to the Order Model or Order Detail Model
//Order Details is inline to Order , Order detail is the list of the items ordered and sumation of prices ordered
type OrderDetail struct {
	uadmin.Model
	ServedTime    *time.Time `uadmin:"filter"`
	TotalDiscount float64    `uadmin:"read_only"`
	GrossSales    float64    `uadmin:"min:0;help:Gross sales are the value of all of Menu Item sumation without accounting for any deductions."`
	NetTotal      float64    `uadmin:"read_only;search;min:0;help:Net Total are gross sales minus deductions: discounts."`
}
