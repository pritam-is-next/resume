package admin

import (
	"fmt"

	"github.com/pritam-is-next/resume/models"
	"github.com/vrianta/agai/v1/controller"
	"github.com/vrianta/agai/v1/log"
)

type Controller struct {
	controller.Context
}

func (c *Controller) GET() controller.View {

	if !c.IsLoggedIn() {
		log.WriteLogf("Unauthorized access to admin panel. Redirecting to login.")
		c.Redirect("/login")
		return controller.EmptyResponse().ToView("login")
	}

	admin_nav_items, err := models.Admin_navItems.Get().Fetch()

	if err != nil {
		fmt.Println("Failed to fetch the admin_navitems")
		return controller.EmptyResponse().ToView("admin")
	}

	user_details, err := models.User_details.Get().First()
	if err != nil {
		fmt.Println("Failed to Fetch User Details")
		return controller.EmptyResponse().ToView("admin")
	}

	// fmt.Println(admin_nav_items)

	response := controller.Response{
		"Title":        "Pritam Dutta",
		"Heading":      "Admin Panel",
		"Name":         "Pritam Dutta",
		"Nav_items":    admin_nav_items,
		"User_Details": user_details,
	}
	// fmt.Println("Admin Panel Accessed")
	return response.ToView("admin")
}
