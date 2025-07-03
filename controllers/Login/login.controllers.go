package login

import (
	"fmt"

	components "github.com/pritam-is-next/resume/Components"
	models "github.com/pritam-is-next/resume/models"
	Controller "github.com/vrianta/Server/controller"
	Log "github.com/vrianta/Server/log"
	Template "github.com/vrianta/Server/template"
)

type login Controller.Struct

var Login = Controller.Struct{
	View: "login",
	GET: func(self *Controller.Struct) *Template.Response {

		if self.IsLoggedIn() {
			self.Redirect("/admin")
			return &Template.EmptyResponse
		}

		response := Template.Response{
			"Title":          "Pritam Dutta",
			"Heading":        "Pritam Dutta",
			"NavItems":       components.NavItems,
			"Hero":           components.Hero,
			"AboutMe":        components.AboutMe,
			"Skills":         components.Skills,
			"Experiences":    components.Experiences,
			"Projects":       components.Projects,
			"ContactDetails": components.ContactDetails,
		}
		return &response
	},
	POST: func(self *Controller.Struct) *Template.Response {

		if self.IsLoggedIn() {
			self.Redirect("/admin")
			Log.WriteLog("Redirecting to Admin")
			return &Template.EmptyResponse
		}

		email := self.GetInput("loginEmail")
		password := self.GetInput("loginPassword")

		if email == nil || password == nil {
			self.Redirect("/login")
			Log.WriteLog("Redirecting to Login")
			return &Template.EmptyResponse
		}

		if users, err := models.Users.Get().
			Where("userName").Is(email.(string)).
			And().
			Where("password").Is(password.(string)).
			Fetch(); err != nil {
			Log.WriteLog("Got error while fetching ", err.Error())
			return &Template.Response{
				"email":    email,
				"password": password,
			}
		} else if len(users) > 0 {
			fmt.Println(users[0].GetFieldValue("userId"))
			self.Login()
			Log.WriteLog("Redirecting to Admin")
			self.Redirect("/admin")
		}

		return &Template.EmptyResponse
	},
}
