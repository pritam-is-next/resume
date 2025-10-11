package home

import (
	components "github.com/pritam-is-next/resume/components"
	models "github.com/pritam-is-next/resume/models"
	"github.com/vrianta/agai/v1"
)

type Controller struct {
	agai.Controller
}

func (c Controller) GET() agai.View {

	initialised, ok := models.App_state.GetComponent("initialised")

	if !ok {
		panic("component is not available")
	}

	val, ok := initialised["Value"].(string)

	if !ok {
		panic("component do not have attrbiute Value")
	}

	if val == "f" {
		// means the application is first time starting and needs to be registered with the user
		c.Redirect("/register")
	}

	nav_items := models.Nav_items.GetComponents()
	response := &agai.Response{
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

	return c.View("home", response)
}
