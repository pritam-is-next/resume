package profile

import (
	"fmt"

	"github.com/pritam-is-next/resume/models"
	"github.com/vrianta/agai/v1"
	"github.com/vrianta/agai/v1/config"
	"github.com/vrianta/agai/v1/log"
)

type Edite struct {
	agai.Controller
}

func (e *Edite) GET() agai.View {
	if !e.IsLoggedIn() {
		log.WriteLogf("Unauthorized profile access. Redirecting to login.\n")
		return e.Redirect("/login")
	}

	uid, uid_ok := e.GetStoredData("uid")
	if !uid_ok {
		log.WriteLog("No User ID Found")
		// log.WriteLog(c.GetStoredDatas())
		return e.Redirect("/admin/profile")
	}
	user_details, err := models.User_details.Get().Where(models.User_details.Fields.UserId).Is(uid).First()
	if err != nil {
		fmt.Println("Failed to fetch user details")
		return e.View("profile", e.EmptyResponse())
	}

	response := agai.Response{
		"Title":        "Profile",
		"Heading":      "Profile",
		"Nav_items":    nil,
		"User_Details": user_details,
	}
	return e.View("edit", response)
}

/*
POST /admin/profile/update
Update profile data
*/
func (c *Edite) POST() agai.View {

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
