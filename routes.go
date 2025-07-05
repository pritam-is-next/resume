package main

import (
	Controllers "github.com/pritam-is-next/resume/controllers"
	Router "github.com/vrianta/Server/v1/router"
)

// Register Routes initializes the routes for the application.
func init() {
	Router.New("/").RegisterRoutes(
		Router.Route("", Controllers.Home),
		Router.Route("home", Controllers.Home),
		Router.Route("admin", Controllers.Admin),
		Router.Route("login", Controllers.Login),
		Router.Route("logout", Controllers.Logout),
	)

	Router.New("/api/").RegisterRoutes()
}
