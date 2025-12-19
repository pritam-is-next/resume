package profile

import (
	"fmt"

	"github.com/pritam-is-next/resume/models"
	"github.com/vrianta/agai/v1"
	"github.com/vrianta/agai/v1/config"
	"github.com/vrianta/agai/v1/log"
)

type Controller struct {
	agai.Controller
}

/*
GET /admin/profile
Show profile page
*/
func (c *Controller) GET() agai.View {

	if !c.IsLoggedIn() {
		log.WriteLogf("Unauthorized profile access. Redirecting to login.\n")
		return c.Redirect("/login")
	}

	admin_nav_items, err := models.Admin_navItems.Get().Fetch()
	if err != nil {
		fmt.Println("Failed to fetch nav items")
		return c.View("profile", c.EmptyResponse())
	}

	user_details, err := models.User_details.Get().First()
	if err != nil {
		fmt.Println("Failed to fetch user details")
		return c.View("profile", c.EmptyResponse())
	}

	response := agai.Response{
		"Title":        "My Profile",
		"Heading":      "Admin Panel",
		"Nav_items":    admin_nav_items,
		"User_Details": user_details,
	}

	return c.View("profile", response)
}

/*
POST /admin/profile/update
Update profile data
*/
func (c *Controller) POST() agai.View {

	if !c.IsLoggedIn() {
		return c.Redirect("/login")
	}

	uid, _ := c.GetStoredData("uid")
	user_details, err := models.User_details.Get().Where(models.Users.Fields.UserName).Is(uid.(string)).First()
	if err != nil {
		fmt.Println("User details not found")
		return c.Redirect("/admin/profile")
	}

	// Update fields
	user_details["FullName"] = c.GetInput("full_name")
	user_details["Email"] = c.GetInput("email")
	user_details["Phone"] = c.GetInput("phone")
	user_details["Dob"] = c.GetInput("dob")
	user_details["Gender"] = c.GetInput("gender")
	user_details["Bio"] = c.GetInput("bio")
	user_details["AddressLine"] = c.GetInput("address_line")
	user_details["City"] = c.GetInput("city")
	user_details["State"] = c.GetInput("state")
	user_details["Country"] = c.GetInput("country")
	user_details["ZipCode"] = c.GetInput("zip_code")

	// Avatar upload (optional)
	if file, ok := c.File("avatar"); ok && file != nil {
		path := config.GetWebConfig().StaticFolders[0] + "/" + file.Filename
		if _, err := c.SaveFile(file, "./"+path); err == nil {
			user_details["Avatar"] = "/static/" + file.Filename
		}
	}

	// Save changes
	models.User_details.Update(models.User_details.Fields.FullName).Exec()
	// if err := user_details.Save(); err != nil {
	// 	fmt.Println("Failed to update profile")
	// 	return c.Redirect("/admin/profile")
	// }

	log.WriteLogf("Profile updated successfully\n")
	return c.Redirect("/admin/profile")
}
