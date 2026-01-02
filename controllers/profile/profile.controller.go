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

	uid, ok := c.GetStoredData("uid")
	if !ok {
		fmt.Println("User ID not found in session")
		return c.Redirect("/login")
	}
	user_details, err := models.User_details.Get().Where(models.Users.Fields.UserName).Is(uid.(string)).First()
	if err != nil {
		fmt.Println("User details not found")
		return c.Redirect("/admin/profile")
	}

	// Update fields
	user_details["FullName"] = c.GetInput("full_name").(string)
	// user_details["Email"] = c.GetInput("email").(string)
	// user_details["Phone"] = c.GetInput("phone").(string)
	// user_details["Dob"] = c.GetInput("dob").(string)
	// user_details["Gender"] = c.GetInput("gender").(string)
	// user_details["Bio"] = c.GetInput("bio").(string)
	// user_details["AddressLine"] = c.GetInput("address_line").(string)
	// user_details["City"] = c.GetInput("city").(string)
	// user_details["State"] = c.GetInput("state").(string)
	// user_details["Country"] = c.GetInput("country").(string)
	// user_details["ZipCode"] = c.GetInput("zip_code").(string)

	// Avatar upload (optional)
	if file, ok := c.File("avatar"); ok && file != nil {
		path := config.GetWebConfig().StaticFolders[0] + "/" + file.Filename
		if _, err := c.SaveFile(file, "./"+path); err == nil {
			user_details["Avatar"] = "/static/" + file.Filename
		}
	}

	// Save changes
	full_name := user_details["FullName"].(string)
	if full_name == "" {
		fmt.Println("Full Name cannot be empty")
		return c.Redirect("/admin/profile")
	}
	models.User_details.Update(models.User_details.Fields.FullName).Where(models.User_details.Fields.FullName).Is(user_details["FullName"]).To(full_name).Exec()
	// models.User_details.Update(models.User_details.Fields.Email).Where(models.User_details.Fields.FullName).Is(user_details["FullName"]).To(user_details["Email"])
	// models.User_details.Update(models.User_details.Fields.Phone).Where(models.User_details.Fields.FullName).Is(user_details["FullName"]).To(user_details["Phone"])
	// models.User_details.Update(models.User_details.Fields.Dob).Where(models.User_details.Fields.FullName).Is(user_details["FullName"]).To(user_details["Dob"])
	// models.User_details.Update(models.User_details.Fields.Gender).Where(models.User_details.Fields.FullName).Is(user_details["FullName"]).To(user_details["Gender"])
	// models.User_details.Update(models.User_details.Fields.Bio).Where(models.User_details.Fields.FullName).Is(user_details["FullName"]).To(user_details["Bio"])
	// models.User_details.Update(models.User_details.Fields.AddressLine).Where(models.User_details.Fields.FullName).Is(user_details["FullName"]).To(user_details["AddressLine"])
	// models.User_details.Update(models.User_details.Fields.City).Where(models.User_details.Fields.FullName).Is(user_details["FullName"]).To(user_details["City"])
	// models.User_details.Update(models.User_details.Fields.State).Where(models.User_details.Fields.FullName).Is(user_details["FullName"]).To(user_details["State"])
	// models.User_details.Update(models.User_details.Fields.Country).Where(models.User_details.Fields.FullName).Is(user_details["FullName"]).To(user_details["Country"])
	// models.User_details.Update(models.User_details.Fields.ZipCode).Where(models.User_details.Fields.FullName).Is(user_details["FullName"]).To(user_details["ZipCode"])
	// models
	// if err := user_details.Save(); err != nil {
	// 	fmt.Println("Failed to update profile")
	// 	return c.Redirect("/admin/profile")
	// }

	log.WriteLogf("Profile updated successfully\n")
	return c.Redirect("/admin/profile")
}
