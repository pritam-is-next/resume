package controllers

import "github.com/vrianta/agai/v1"

type Controller struct {
	agai.Controller
}

func (c *Controller) GET() agai.View {
	return c.Redirect("/home")
}
