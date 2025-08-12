package main

import (
	Controllers "github.com/pritam-is-next/resume/controllers"
	"github.com/pritam-is-next/resume/controllers/register"
	"github.com/pritam-is-next/resume/controllers/test"
	Router "github.com/vrianta/agai/v1/router"
)

// Register Routes initializes the routes for the application.
func init() {
	Router.New("/").RegisterRoutes(
		Router.Route("", Controllers.Home),
		Router.Route("home", Controllers.Home),
		Router.Route("admin", Controllers.Admin),
		Router.Route("login", Controllers.Login),
		Router.Route("logout", Controllers.Logout),
		Router.Route("register", register.Controller),
		Router.Route("test", test.Controller),
	)

	Router.New("/api/").RegisterRoutes()
}
