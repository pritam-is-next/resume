package login

import (
	components "github.com/pritam-is-next/resume/components"
	models "github.com/pritam-is-next/resume/models"
	Controller "github.com/vrianta/agai/v1/controller"
	Log "github.com/vrianta/agai/v1/log"
	Template "github.com/vrianta/agai/v1/template"
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

		if user, err := models.Users.Get().
			Where("userName").Is(email.(string)).
			And().
			Where("password").Is(password.(string)).
			First(); err != nil {
			Log.WriteLog("Got error while fetching ", err.Error())
			return &Template.Response{
				"email":    email,
				"password": password,
			}
		} else if user != nil {
			// fmt.Println(user["userId"])
			self.Login()
			Log.WriteLog("Redirecting to Admin")
			self.Redirect("/admin")
		}

		return &Template.EmptyResponse
	},
}
