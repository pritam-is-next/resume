package home

import (
	"fmt"

	components "github.com/pritam-is-next/resume/components"
	models "github.com/pritam-is-next/resume/models"
	"github.com/vrianta/agai/v1"
	"github.com/vrianta/agai/v1/log"
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
	user_details, ud_err := models.User_details.Get().First()
	if ud_err != nil {
		log.Error("Not able to fetch userdetails %s", ud_err.Error())
	}
	response := &agai.Response{
		"Title":          "Pritam Dutta",
		"Heading":        "Pritam Dutta",
		"NavItems":       nav_items,
		"Hero":           components.Hero,
		"AbouteMe":       user_details["AbouteMe"],
		"Avatar":         user_details["Avatar"],
		"Skills":         components.Skills,
		"Experiences":    components.Experiences,
		"Projects":       components.Projects,
		"ContactDetails": components.ContactDetails,
	}

	if settings, ok := models.Settings.GetComponent("Theme"); !ok {
		log.Error("Setting theme not found")
	} else {
		fmt.Println(settings)
		if theme_name, ok := settings.FieldValue("Value"); !ok {
			log.Error("Theme name not found")
		} else {
			log.WriteLog("Theme Name: ", theme_name)
			return c.View(theme_name.(string)+"/home", response)
		}
	}

	return c.View("home", response)
}
