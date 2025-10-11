package logout

import (
	"github.com/vrianta/agai/v1"
)

type Controller struct {
	agai.Controller
}

func (c *Controller) GET() agai.View {
	c.Logout()
	c.Redirect("/")
	return agai.EmptyResponse().AsView("home")
}
