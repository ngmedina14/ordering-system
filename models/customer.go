/*
Welcome to the Ordering System Documentation

Customer Model

this is to collect their data and use it as identifier
hello world
*/
package models

import (
	"github.com/uadmin/uadmin"
)

//Customer is the person ordering before entering the Menu Category
type Customer struct {
	uadmin.Model
	GuestName   string `uadmin:"search" sql:"type:text;"`
	MacAddress  string `uadmin:"search" sql:"type:text;"`
	FacebookAPI string `sql:"type:text;"`
}

//String function used as Identifier
func (c *Customer) String() string {
	return c.GuestName
}
