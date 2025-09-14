package admin

import (
	"fmt"

	"github.com/pritam-is-next/resume/models"
	Controller "github.com/vrianta/agai/v1/controller"
	Template "github.com/vrianta/agai/v1/template"
)

var Admin = Controller.Context{
	View: "admin",
	GET:  GET,
}

var GET = func(self *Controller.Context) *Template.Response {

	if !self.IsLoggedIn() {
		self.Redirect("/login")
	}

	admin_nav_items, err := models.Admin_navItems.Get().Fetch()

	if err != nil {
		fmt.Println("Failed to fetch the admin_navitems")
		return &Template.NoResponse
	}

	user_details, err := models.User_details.Get().First()
	if err != nil {
		fmt.Println("Failed to Fetch User Details")
		return &Template.NoResponse
	}

	fmt.Println(admin_nav_items)

	response := Template.Response{
		"Title":        "Pritam Dutta",
		"Heading":      "Admin Panel",
		"Name":         "Pritam Dutta",
		"Nav_items":    admin_nav_items,
		"User_Details": user_details,
	}

	return &response
}
