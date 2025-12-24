package main

import (
	"github.com/pritam-is-next/resume/controllers"
	"github.com/pritam-is-next/resume/controllers/admin"
	"github.com/pritam-is-next/resume/controllers/admin/profile"
	"github.com/pritam-is-next/resume/controllers/home"
	"github.com/pritam-is-next/resume/controllers/login"
	loginapi "github.com/pritam-is-next/resume/controllers/loginapi"
	"github.com/pritam-is-next/resume/controllers/logout"
	"github.com/pritam-is-next/resume/controllers/register"
	"github.com/vrianta/agai/v1"
)

// Register Routes initializes the routes for the application.
func init() {
	agai.CreateRoute[controllers.Controller]("/")
	agai.CreateRoute[home.Controller]("/home")
	agai.CreateRoute[admin.Controller]("/admin")
	agai.CreateRoute[login.Controller]("/login")
	agai.CreateRoute[loginapi.Controller]("/api/login")
	agai.CreateRoute[logout.Controller]("/logout")
	agai.CreateRoute[register.Controller]("/register")
	agai.CreateRoute[profile.Controller]("/admin/profile")
	agai.CreateRoute[profile.Edite]("/admin/profile/edit")
}
