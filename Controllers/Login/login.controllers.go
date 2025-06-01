package login

import (
	components "github.com/pritam-is-next/resume/Components"
	"github.com/vrianta/Server/Controller"
	"github.com/vrianta/Server/Log"
	"github.com/vrianta/Server/Redirect"
	"github.com/vrianta/Server/Template"
)

type login Controller.Struct

var Login = Controller.Struct{
	View: "login.php",
	GET: func(self *Controller.Struct) *Template.Response {

		if self.Session.IsLoggedIn() {
			Redirect.New("/admin", self.Session).Redirect()
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

		if self.Session.IsLoggedIn() {
			Redirect.New("/admin", self.Session).Redirect()
			return &Template.EmptyResponse
		}

		Log.WriteLog(self.Session.POST)
		email, email_ok := self.Session.POST["loginEmail"]
		password, password_ok := self.Session.POST["loginPassword"]

		if !email_ok || !password_ok {
			Redirect.New("", self.Session).Redirect()
		}

		if email == "sample@gmail.com" && password == "pass" {
			self.Session.Login("sample@gmail.com")
			Redirect.New("/admin", self.Session).Redirect()
		}

		return &Template.EmptyResponse
	},
}
