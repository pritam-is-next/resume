package admin

import (
	components "github.com/pritam-is-next/resume/components"
	"github.com/pritam-is-next/resume/models"
	Controller "github.com/vrianta/agai/v1/controller"
	Template "github.com/vrianta/agai/v1/template"
)

var Admin = Controller.Context{
	View: "admin",
	GET:  GET,
}

var GET = func(self *Controller.Context) *Template.Response {

	if !self.GetSession().IsLoggedIn() {
		self.Redirect("/login")
	}

	response := Template.Response{
		"Title":          "Pritam Dutta",
		"Heading":        "Pritam Dutta",
		"NavItems":       models.Nav_items.GetComponents(),
		"Hero":           components.Hero,
		"AboutMe":        components.AboutMe,
		"Skills":         components.Skills,
		"Experiences":    components.Experiences,
		"Projects":       components.Projects,
		"ContactDetails": components.ContactDetails,
	}

	return &response
}
