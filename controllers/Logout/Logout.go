package logout

import (
	Controller "github.com/vrianta/agai/v1/controller"
	Template "github.com/vrianta/agai/v1/template"
)

var Logout = Controller.Struct{
	GET: func(self *Controller.Struct) *Template.Response {
		self.Logout()
		self.Redirect("/")
		return &Template.NoResponse
	},
}
