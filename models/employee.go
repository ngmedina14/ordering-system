/*
Employee model

Employee model collects the data of the staff

also this model provide them access to the system

and used to create cash drawer
*/
package models

import "github.com/uadmin/uadmin"

type Employee struct {
	uadmin.Model
	User                 uadmin.User `uadmin:"help:Select the user in the choices.;list_exclude;hidden"`
	UserID               uint
	FirstName            string `uadmin:"search"`
	LastName             string `uadmin:"search"`
	EmployeeID           string `uadmin:"search"`
	Mobile               string `uadmin:"help:Employee's cellphone number;list_exclude"`
	Passcode             string `uadmin:"read_only;help:System-generated;list_exclude"`
	UserName             string `uadmin:"required"`
	Password             string `uadmin:"list_exclude"`
	DisableConsoleAccess bool   `uadmin:"help:Remove access on Console"`
}

func (e *Employee) String() string {
	return e.User.String()
}
