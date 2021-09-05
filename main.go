/*
This is an Ordering System
referenced to sharecafe and wendys design

collaboration of shanelle harold and romnick

Please dont get to excited
*/
package main

import (
	"net/http"

	"github.com/ngmedina14/OrderingSystem/models"
	"github.com/ngmedina14/OrderingSystem/views"
	"github.com/uadmin/uadmin"
)

func main() {
	DBConfig()
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
	uadmin.RootURL = "/admin/"
	// http.HandleFunc("/console/", uadmin.Handler(views.ConsoleHandler))
	// http.HandleFunc("/console/login/", uadmin.Handler(views.LoginHandler))
	http.HandleFunc("/", uadmin.Handler(views.PageHandler))
	// http.HandleFunc("/app-v2/", uadmin.Handler(views.AppHandlerV2))
	// http.HandleFunc("/api/", uadmin.Handler(api.APIHandler))
	// http.HandleFunc("/pos/", uadmin.Handler(views.POSHandler))
	// http.HandleFunc("/pos/login", uadmin.Handler(views.POSLoginHandler))
	uadmin.StartServer()
}

// DBConfig !
func DBConfig() {
	// Change DB Setting to MySQL
	uadmin.Database = &uadmin.DBSettings{
		Type: "mysql",
		Host: "orderingsystem_mysql_1",
		// Host:     "192.168.150.13",
		Name:     "ordering-system",
		User:     "root",
		Password: "NeilGwapo100%",
		Port:     3306,
	}
}
