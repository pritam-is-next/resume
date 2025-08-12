package register

import (
	"github.com/pritam-is-next/resume/models"
	"github.com/vrianta/agai/v1/controller"
	"github.com/vrianta/agai/v1/template"
)

/*
Register Controller

This is a basic controller built on top of `controller.Context`. It supports all the usual HTTP methods:
GET, POST, PUT, DELETE, PATCH, HEAD, and OPTIONS.

Each method is just a variable on the struct — you can set any of them to a handler function like:

	func(self *controller.Context) *template.Response

If you don't define a method, GET will be used by default. So if your controller just needs to respond to GET requests,
you only need to set that and you're good to go.

This setup keeps things simple — define what you need, skip what you don't.
*/

var Controller = controller.Context{
	View: "register",
	GET:  GET,
}

var GET = func(self *controller.Context) *template.Response {

	if user, err := models.Users.Get().First(); err == nil {
		// that means it is a initial application loading so we have to enable the resigtration page
		// else we can move directly to the home page or loading page

		if user.IsValid() {
			self.Redirect("/")
		}
	}

	// means it is a fresh application start it needs to register the first user
	response := &template.Response{
		"controller_name": "Register",
	}

	return response
}
