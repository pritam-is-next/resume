package Controllers

import (
	components "github.com/pritam-is-next/resume/Components"
	Server "github.com/vrianta/Server"
)

type Home struct{}

func (h *Home) GET(Session *Server.Session) {
	if err := Session.RenderEngine.RenderTemplate("home.php", Server.TemplateData{
		"Title":       "Pritam Dutta",
		"Heading":     "Pritam Dutta",
		"NavItems":    components.NavItems,
		"Hero":        components.Hero,
		"AboutMe":     components.AboutMe,
		"Skills":      components.Skills,
		"Experiences": components.Experiences,
		"Projects":    components.Projects,
	}); err != nil {
		Session.RenderEngine.Render(err.Error())
	}
}

func (h *Home) POST(Session *Server.Session) {

}

func (h *Home) DELETE(Session *Server.Session) {

}
