package logout

import (
	"github.com/vrianta/Server/Controller"
	"github.com/vrianta/Server/Template"
)

var Logout = Controller.Struct{
	GET: func(self *Controller.Struct) *Template.Response {
		self.Logout()
		self.Redirect("/")
		return &Template.NoResponse
	},
}
