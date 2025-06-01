package admin

import (
	components "github.com/pritam-is-next/resume/Components"
	Controller "github.com/vrianta/Server/Controller"
	"github.com/vrianta/Server/Redirect"
	Template "github.com/vrianta/Server/Template"
)

var Admin = Controller.Struct{
	View: "admin.php",
	GET:  GET,
}

var GET = func(self *Controller.Struct) *Template.Response {

	if !self.Session.IsLoggedIn() {
		Redirect.New("/login", self.Session).Redirect()
	}

	response := Template.Response{
		"Title":          "Pritam Dutta",
		"Heading":        "Pritam Dutta",
		"NavItems":       components.NavItems,
		"Hero":           components.Hero,
		"AboutMe":        components.AboutMe,
		"Skills":         components.Skills,
		"Experiences":    components.Experiences,
		"Projects":       components.Projects,
		"ContactDetails": components.ContactDetails,
	}

	return &response
}
