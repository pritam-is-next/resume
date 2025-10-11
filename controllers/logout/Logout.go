package logout

import (
	"github.com/vrianta/agai/v1"
)

type Controller struct {
	agai.Controller
}

func (c *Controller) GET() agai.View {
	c.Logout()
	return c.Redirect("/")
}
