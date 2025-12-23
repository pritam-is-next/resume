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

	uid, uid_ok := c.GetStoredData("uid")
	if !uid_ok {
		log.WriteLog("No User ID Found")
		// log.WriteLog(c.GetStoredDatas())
		return c.Redirect("/admin/profile")
	}
	user_details, err := models.User_details.Get().Where(models.User_details.Fields.UserId).Is(uid).First()
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

	uid, uid_ok := c.GetStoredData("uid")
	if !uid_ok {
		log.WriteLog("No User ID Found")
		log.WriteLog(c.GetStoredDatas())
		return c.Redirect("/admin/profile")
	}
	user_details, err := models.User_details.Get().Where(models.Users.Fields.UserId).Is(uid.(string)).First()
	if err != nil {
		fmt.Println("User details not found ", err.Error())
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
			user_details["Avatar"] = "/" + path
		}
	}

	query := models.User_details.
		Update(nil).
		Where(models.User_details.Fields.UserId).
		Is(uid)

	query.
		Set(models.User_details.Fields.FullName).To(user_details["FullName"]).
		Set(models.User_details.Fields.Email).To(user_details["Email"]).
		Set(models.User_details.Fields.Phone).To(user_details["Phone"]).
		Set(models.User_details.Fields.Dob).To(user_details["Dob"]).
		Set(models.User_details.Fields.Gender).To(user_details["Gender"]).
		Set(models.User_details.Fields.Bio).To(user_details["Bio"]).
		Set(models.User_details.Fields.AddressLine).To(user_details["AddressLine"]).
		Set(models.User_details.Fields.City).To(user_details["City"]).
		Set(models.User_details.Fields.State).To(user_details["State"]).
		Set(models.User_details.Fields.Country).To(user_details["Country"]).
		Set(models.User_details.Fields.ZipCode).To(user_details["ZipCode"])

	// Only update avatar if it exists
	if avatar, ok := user_details["Avatar"]; ok {
		query.Set(models.User_details.Fields.Avatar).To(avatar)
	}

	// Execute
	if err := query.Exec(); err != nil {
		log.Error("failed to execute query %s", err.Error())
		return c.Redirect("/admin/profile")
	}

	log.WriteLogf("Profile updated successfully\n")
	return c.Redirect("/admin/profile")
}
