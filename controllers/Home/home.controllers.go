package home

import (
	components "github.com/pritam-is-next/resume/components"
	models "github.com/pritam-is-next/resume/models"
	Controller "github.com/vrianta/agai/v1/controller"
	Template "github.com/vrianta/agai/v1/template"
)

var Home = Controller.Struct{
	View: "Home",
	GET:  GET,
}

var GET = func(self *Controller.Struct) *Template.Response {
	nav_items := models.Nav_items.GetComponents()
	response := &Template.Response{
		"Title":          "Pritam Dutta",
		"Heading":        "Pritam Dutta",
		"NavItems":       nav_items,
		"Hero":           components.Hero,
		"AboutMe":        components.AboutMe,
		"Skills":         components.Skills,
		"Experiences":    components.Experiences,
		"Projects":       components.Projects,
		"ContactDetails": components.ContactDetails,
	}

	return response
}
