package main

import (
	"github.com/ngmedina14/OrderingSystem/models"
	"github.com/uadmin/uadmin"
)

func main() {

	uadmin.Register(
		models.ComboItem{},
		models.Customer{},
		models.Discount{},
		models.Employee{},
		models.ItemOption{},
		models.MenuCategory{},
		models.OfficialReceipt{},
		models.OrderDetail{},
		models.Order{},
		models.Payment{},
		models.Tax{},
	)

	uadmin.StartServer()
}
