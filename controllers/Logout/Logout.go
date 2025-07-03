package logout

import (
	Controller "github.com/vrianta/Server/controller"
	Template "github.com/vrianta/Server/template"
)

var Logout = Controller.Struct{
	GET: func(self *Controller.Struct) *Template.Response {
		self.Logout()
		self.Redirect("/")
		return &Template.NoResponse
	},
}
