/*
This is an Ordering System
referenced to sharecafe and wendys design

collaboration of shanelle harold and romnick

Please dont get to excited
*/
package main

import (
	"net/http"
	"os"

	"github.com/ngmedina14/ordering-system/models"
	"github.com/ngmedina14/ordering-system/views"
	"github.com/uadmin/uadmin"
)

func main() {
	DBConfig()
	RegisterModels()
	RegisterInlines()
	PageHandlers()
	WebSockets()
	CustomModelData()
	PageSettings()
}

// DBConfig !
func DBConfig() {
	// Change DB Setting to MySQL
	uadmin.Database = &uadmin.DBSettings{
		Type:     "mysql",
		Host:     os.Getenv("MYSQL_ROOT_HOST"),
		Name:     os.Getenv("MYSQL_DATABASE"),
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_ROOT_PASSWORD"),
		Port:     3306,
	}
}

func RegisterModels() {
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
}
func RegisterInlines() {

}
func PageHandlers() {
	// http.HandleFunc("/console/", uadmin.Handler(views.ConsoleHandler))
	// http.HandleFunc("/console/login/", uadmin.Handler(views.LoginHandler))
	http.HandleFunc("/", uadmin.Handler(views.PageHandler))
	// http.HandleFunc("/app-v2/", uadmin.Handler(views.AppHandlerV2))
	// http.HandleFunc("/api/", uadmin.Handler(api.APIHandler))
	// http.HandleFunc("/pos/", uadmin.Handler(views.POSHandler))
	// http.HandleFunc("/pos/login", uadmin.Handler(views.POSLoginHandler))
}
func WebSockets() {

}
func CustomModelData() {

}

func PageSettings() {
	uadmin.Port = 8080
	uadmin.RootURL = "/admin/"
	uadmin.SiteName = "Ordering System"

	uadmin.StartServer()
}
