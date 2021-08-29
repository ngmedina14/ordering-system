package uadmin

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

// forgotPasswordHandler !
func forgotPasswordHandler(u *User, r *http.Request) error {
	if u.Email == "" {
		return fmt.Errorf("unable to reset password, the user does not have an email")
	}
	msg := `Dear {NAME},

Have you forgotten your password to access {WEBSITE}. Don't worry we got your back. Please follow the link below to reset your password.

If you want to reset your password, click this link:
<a href="{URL}">{URL}</a>

If you didn't request a password reset, you can ignore this message.

Regards,
{WEBSITE} Support
`
	// Check if the host name is in the allowed hosts list
	allowed := false
	var host string
	var allowedHost string
	var err error
	if host, _, err = net.SplitHostPort(r.Host); err != nil {
		host = r.Host
	}
	for _, v := range strings.Split(AllowedHosts, ",") {
		if allowedHost, _, err = net.SplitHostPort(v); err != nil {
			allowedHost = v
		}
		if allowedHost == host {
			allowed = true
		}
	}
	if !allowed {
		Trail(CRITICAL, "Reset password request for host: (%s) which is not in AllowedHosts settings", host)
		return nil
	}

	urlParts := strings.Split(r.Header.Get("origin"), "://")
	link := urlParts[0] + "://" + r.Host + RootURL + "resetpassword?u=" + fmt.Sprint(u.ID) + "&key=" + u.GetOTP()
	msg = strings.Replace(msg, "{NAME}", u.String(), -1)
	msg = strings.Replace(msg, "{WEBSITE}", SiteName, -1)
	msg = strings.Replace(msg, "{URL}", link, -1)
	subject := "Password reset for " + SiteName
	err = SendEmail([]string{u.Email}, []string{}, []string{}, subject, msg)

	return err
}
