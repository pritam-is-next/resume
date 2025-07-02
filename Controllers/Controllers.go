package controllers

import (
	admin "github.com/pritam-is-next/resume/Controllers/admin"
	home "github.com/pritam-is-next/resume/Controllers/home"
	login "github.com/pritam-is-next/resume/Controllers/login"
	logout "github.com/pritam-is-next/resume/Controllers/logout"
)

var (
	Home   = home.Home
	Admin  = admin.Admin
	Login  = login.Login
	Logout = logout.Logout
)
