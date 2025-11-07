package loginapi

import "github.com/vrianta/agai/v1"

type Controller struct {
	agai.Controller
}

func (c *Controller) GET() agai.View {
	if c.IsLoggedIn() {
		return c.Redirect("/admin")
	}
	return c.Redirect("/")
}
