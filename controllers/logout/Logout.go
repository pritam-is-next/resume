package logout

import (
	"github.com/vrianta/agai/v1/controller"
)

type Controller struct {
	controller.Context
}

func (c *Controller) GET() controller.View {
	c.Logout()
	c.Redirect("/")
	return controller.EmptyResponse().ToView("home")
}
