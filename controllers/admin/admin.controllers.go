package admin

import (
	"fmt"

	"github.com/pritam-is-next/resume/models"
	"github.com/vrianta/agai/v1"
	"github.com/vrianta/agai/v1/log"
)

type Controller struct {
	agai.Controller
}

func (c *Controller) GET() agai.View {

	if !c.IsLoggedIn() {
		log.WriteLogf("Unauthorized access to admin panel. Redirecting to login.\n")
		return c.Redirect("/login")
	}

	admin_nav_items, err := models.Admin_navItems.Get().Fetch()

	if err != nil {
		fmt.Println("Failed to fetch the admin_navitems")
		return c.View("admin", c.EmptyResponse())
	}

	user_details, err := models.User_details.Get().First()
	if err != nil {
		fmt.Println("Failed to Fetch User Details")
		return c.View("admin", c.EmptyResponse())
	}

	// fmt.Println(admin_nav_items)

	response := agai.Response{
		"Title":        "Pritam Dutta",
		"Heading":      "Admin Panel",
		"Name":         "Pritam Dutta",
		"Nav_items":    admin_nav_items,
		"User_Details": user_details,
	}
	// fmt.Println("Admin Panel Accessed")
	return c.View("admin", response)
}
