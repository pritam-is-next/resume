package login

import (
	components "github.com/pritam-is-next/resume/Components"
	"github.com/vrianta/Server/Controller"
	"github.com/vrianta/Server/Log"
	"github.com/vrianta/Server/Template"
)

type login Controller.Struct

var Login = Controller.Struct{
	View: "login.php",
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

		if email.(string) == "sample@gmail.com" && password.(string) == "pass" {
			self.Login()
			Log.WriteLog("Redirecting to Admin")
			self.Redirect("/admin")
		}

		return &Template.EmptyResponse
	},
}
