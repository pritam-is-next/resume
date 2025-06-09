package main

import (
	"github.com/pritam-is-next/resume/Controllers"
	"github.com/vrianta/Server/Router"
)

var rootPths = Router.New("/").RegisterRoutes(
	Router.Route("", Controllers.Home),
	Router.Route("admin", Controllers.Admin),
	Router.Route("login", Controllers.Login),
	Router.Route("logout", Controllers.Logout),
)

var apiRoutes = Router.New("/api/").RegisterRoutes()
