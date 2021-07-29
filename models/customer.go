package models

import "github.com/uadmin/uadmin"

//Customer is the person ordering before entering the Menu Category
type Customer struct {
	uadmin.Model
	GuestName   string `uadmin:"search" sql:"type:text;"`
	MacAddress  string `uadmin:"search" sql:"type:text;"`
	FacebookAPI string `sql:"type:text;"`
}

func (c *Customer) String() string {
	return c.GuestName
}
