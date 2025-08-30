package home

import (
	components "github.com/pritam-is-next/resume/components"
	models "github.com/pritam-is-next/resume/models"
	Controller "github.com/vrianta/agai/v1/controller"
	Template "github.com/vrianta/agai/v1/template"
)

var Home = Controller.Context{
	View: "home",
	GET:  GET,
}

var GET = func(self *Controller.Context) *Template.Response {

	initialised, ok := models.App_state.GetComponent("initialised")

	if !ok {
		panic("component is not available")
	}

	val, ok := initialised["Value"].(string)

	if !ok {
		panic("component do not have attrbiute Value")
	}

	if val == "false" {
		// means the application is first time starting and needs to be registered with the user
		self.Redirect("/register")
	}

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
