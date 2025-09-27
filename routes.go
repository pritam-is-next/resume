package main

import (
	"github.com/pritam-is-next/resume/controllers/admin"
	"github.com/pritam-is-next/resume/controllers/home"
	"github.com/pritam-is-next/resume/controllers/login"
	"github.com/pritam-is-next/resume/controllers/logout"
	"github.com/pritam-is-next/resume/controllers/profile"
	"github.com/pritam-is-next/resume/controllers/register"
	"github.com/vrianta/agai/v1/router"
)

// Register Routes initializes the routes for the application.
func init() {
	// Router.CreateRoute("/home", &home.Controller{})
	// Router.CreateRoute("/admin", &admin.Controller{})
	// Router.CreateRoute("/login", &login.Controller{})
	// Router.CreateRoute("/logout", &logout.Controller{})
	// Router.CreateRoute("/register", &register.Controller{})
	// Router.CreateRoute("/admin/profile", &profile.Controller{})

	router.New("").
		AddRoute("home", &home.Controller{}).
		AddRoute("admin", &admin.Controller{}).
		AddRoute("login", &login.Controller{}).
		AddRoute("logout", &logout.Controller{}).
		AddRoute("register", &register.Controller{}).
		AddRoute("admin/profile", &profile.Controller{})
	// Router.New("/").RegisterRoutes(
	// 	Router.Route("", Controllers.Home),
	// 	Router.Route("home", Controllers.Home),
	// 	Router.Route("admin", Controllers.Admin),
	// 	Router.Route("login", Controllers.Login),
	// 	Router.Route("logout", Controllers.Logout),
	// 	Router.Route("register", register.Controller),
	// 	Router.Route("test", test.Controller),
	// 	Router.Route("admin/profile", profile.Controller),
	// )

	// Router.New("/api/").RegisterRoutes()
}
