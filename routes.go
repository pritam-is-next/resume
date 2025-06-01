package main

import (
	"github.com/pritam-is-next/resume/Controllers"
	"github.com/vrianta/Server/Router"
)

var Routes = Router.InitRoutes(&Router.Routes{
	"/":       Controllers.Home,
	"/admin":  Controllers.Admin,
	"/login":  Controllers.Login,
	"/logout": Controllers.Logout,
})
