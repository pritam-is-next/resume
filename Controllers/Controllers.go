package Controllers

import (
	admin "github.com/pritam-is-next/resume/Controllers/Admin"
	home "github.com/pritam-is-next/resume/Controllers/Home"
	login "github.com/pritam-is-next/resume/Controllers/Login"
	logout "github.com/pritam-is-next/resume/Controllers/Logout"
)

var (
	Home   = home.Home
	Admin  = admin.Admin
	Login  = login.Login
	Logout = logout.Logout
)
