package main

import (
	"github.com/pritam-is-next/resume/Controllers"
	Server "github.com/vrianta/Server"
)

var Routes = Server.Routes{
	"/":      Controllers.Home,
	"/admin": Controllers.Admin,
	"/login": Controllers.Login,
}
