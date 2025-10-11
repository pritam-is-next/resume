package main

import (
	"github.com/pritam-is-next/resume/controllers/admin"
	"github.com/pritam-is-next/resume/controllers/home"
	"github.com/pritam-is-next/resume/controllers/login"
	"github.com/pritam-is-next/resume/controllers/logout"
	"github.com/pritam-is-next/resume/controllers/profile"
	"github.com/pritam-is-next/resume/controllers/register"
	"github.com/vrianta/agai/v1"
)

// Register Routes initializes the routes for the application.
func init() {
	agai.CreateRoute[home.Controller]("/")
	agai.CreateRoute[home.Controller]("/home")
	agai.CreateRoute[admin.Controller]("/admin")
	agai.CreateRoute[login.Controller]("/login")
	agai.CreateRoute[logout.Controller]("/logout")
	agai.CreateRoute[register.Controller]("/register")
	agai.CreateRoute[profile.Controller]("/admin/profile")

	// router.New("").
	// 	AddRoute("", &home.Controller{}).
	// 	AddRoute("home", &home.Controller{}).
	// 	AddRoute("admin", &admin.Controller{}).
	// 	AddRoute("login", &login.Controller{}).
	// 	AddRoute("logout", &logout.Controller{}).
	// 	AddRoute("register", &register.Controller{}).
	// 	AddRoute("admin/profile", &profile.Controller{})
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
