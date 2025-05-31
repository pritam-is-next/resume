package main

import (
	Controllers "github.com/pritam-is-next/resume/Controllers/Home"
	server "github.com/vrianta/Server"
)

var Routes = server.Routes{
	"/": Controllers.Home,
	// "/admin": &Controllers.Admin{},
	// "/login": &Controllers.Login{},
}
