package views

import (
	"net/http"

	"github.com/uadmin/uadmin"
)

func PageHandler(w http.ResponseWriter, r *http.Request) {
	// if r.Host == models.Domain {
	// 	http.Redirect(w, r, "http://"+models.Domain+"/console", 303)
	// 	return
	// }
	uadmin.RenderHTML(w, r, "./templates/console/console_login.html", nil)
}
