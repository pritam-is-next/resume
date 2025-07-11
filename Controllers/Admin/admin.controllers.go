package admin

import (
	components "github.com/pritam-is-next/resume/Components"
	Controller "github.com/vrianta/Server/Controller"
	Template "github.com/vrianta/Server/Template"
)

var Admin = Controller.Struct{
	View: "admin",
	GET:  GET,
}

var GET = func(self *Controller.Struct) *Template.Response {

	if !self.GetSession().IsLoggedIn() {
		self.Redirect("/login")
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
