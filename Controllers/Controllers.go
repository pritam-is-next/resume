package controllers

import (
	admin "github.com/pritam-is-next/resume/controllers/admin"
	home "github.com/pritam-is-next/resume/controllers/home"
	login "github.com/pritam-is-next/resume/controllers/login"
	logout "github.com/pritam-is-next/resume/controllers/logout"
)

var (
	Home   = home.Home
	Admin  = admin.Admin
	Login  = login.Login
	Logout = logout.Logout
)
