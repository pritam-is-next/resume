package test

import (
	"github.com/vrianta/agai/v1/controller"
	"github.com/vrianta/agai/v1/template"
)

var Controller = controller.Context{
	View: "test",
	GET: func(self *controller.Context) *template.Response {
		return &template.Response{
			"controller_name": "Test",
		}
	},
}
