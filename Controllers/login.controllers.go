package Controllers

import (
	components "github.com/pritam-is-next/resume/Components"
	Server "github.com/vrianta/Server"
)

type Login struct{}

func (h *Login) GET(Session *Server.Session) {

	if Session.IsLoggedIn() {
		Session.Redirect("/admin")
	}
	if err := Session.RenderEngine.RenderTemplate("login.php", Server.TemplateData{
		"Title":          "Pritam Dutta",
		"Heading":        "Pritam Dutta",
		"NavItems":       components.NavItems,
		"Hero":           components.Hero,
		"AboutMe":        components.AboutMe,
		"Skills":         components.Skills,
		"Experiences":    components.Experiences,
		"Projects":       components.Projects,
		"ContactDetails": components.ContactDetails,
	}); err != nil {
		Session.RenderEngine.Render(err.Error())
	}
}

func (h *Login) POST(Session *Server.Session) {

	if Session.IsLoggedIn() {
		Session.Redirect("/admin")
	}

	email, email_ok := Session.POST["loginEmail"]
	password, password_ok := Session.POST["loginPassword"]

	if !email_ok || !password_ok {
		Session.Redirect("")
	}

	if email == "sample@gmail.com" && password == "pass" {
		Session.Login("sample@gmail.com")
	}
}

func (h *Login) DELETE(Session *Server.Session) {

}
