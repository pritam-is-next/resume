package Controllers

import (
	admin "github.com/pritam-is-next/resume/Controllers/Admin"
	home "github.com/pritam-is-next/resume/Controllers/Home"
	login "github.com/pritam-is-next/resume/Controllers/Login"
)

var (
	Home  = &home.Home
	Admin = &admin.Admin
	Login = &login.Login
)
