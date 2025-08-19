package login

import (
	components "github.com/pritam-is-next/resume/components"
	models "github.com/pritam-is-next/resume/models"
	Controller "github.com/vrianta/agai/v1/controller"
	Log "github.com/vrianta/agai/v1/log"
	"github.com/vrianta/agai/v1/template"
	"github.com/vrianta/agai/v1/utils"
)

var Login = Controller.Context{
	View: "login",
	GET: func(self *Controller.Context) *template.Response {

		if self.IsLoggedIn() {
			self.Redirect("/admin")
			return &template.EmptyResponse
		}

		response := template.Response{
			"Title":          "Pritam Dutta",
			"Heading":        "Pritam Dutta",
			"NavItems":       models.Nav_items.GetComponents(),
			"Hero":           components.Hero,
			"AboutMe":        components.AboutMe,
			"Skills":         components.Skills,
			"Experiences":    components.Experiences,
			"Projects":       components.Projects,
			"ContactDetails": components.ContactDetails,
		}
		return &response
	},
	POST: func(self *Controller.Context) *template.Response {

		if self.IsLoggedIn() {
			self.Redirect("/admin")
			Log.WriteLog("Redirecting to Admin")
			return &template.EmptyResponse
		}

		email := self.GetInput("loginEmail")
		password := self.GetInput("loginPassword")

		if email == nil || password == nil {
			return &template.Response{
				"error": "email or password field is empty",
			}
		}

		if hashed_password, err := utils.HashPassword(password.(string)); err != nil {
			return &template.Response{
				"error": "Internal server error | failed to hash password",
			}

		} else {
			if user, err := models.Users.Get().
				Where("UserName").Is(email.(string)).
				And().
				Where("Password").Is(hashed_password).
				First(); err != nil {
				Log.WriteLog("Got error while fetching ", err.Error())
				return &template.Response{
					"email":    email,
					"password": password,
					"error":    err.Error(),
				}
			} else if user != nil {
				self.Login()
				// Log.WriteLog("Redirecting to Admin")
				self.Redirect("/admin")
			} else {
				return &template.Response{
					"UserName": email,
					"Password": password,
					"error":    "User Name or Password is wrong",
				}
			}
		}

		return &template.EmptyResponse
	},
}
