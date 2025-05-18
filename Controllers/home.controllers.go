package Controllers

import (
	components "github.com/pritam-is-next/resume/Components"
	Server "github.com/pritam-is-next/resume/Server"
)

type Home struct{}

func (h *Home) GET(Session *Server.Session) {
	if err := Session.RenderEngine.RenderTemplate("home.php", Server.TemplateData{
		"Title":   "Pritam Dutta",
		"Heading": "Pritam Dutta",
		"NavItems": []components.NavItems{
			{Name: "Home", Href: "#home", Disabled: false},
			{Name: "About Me", Href: "#about-me", Disabled: false},
			{Name: "Skills", Href: "#skills", Disabled: false},
			{Name: "Experience", Href: "#experience", Disabled: false},
			{Name: "Projects", Href: "#projects", Disabled: false},
			{Name: "Blogs", Href: "/blogs", Disabled: false},
			{Name: "DropDown", Href: "#", Disabled: false, IsDropDown: true, DropDown: []components.DropDown{
				{Name: "Dropdown1", Href: "#"},
				{Name: "Dropdown2", Href: "#"},
			}},
		},
	}); err != nil {
		Session.RenderEngine.Render(err.Error())
	}
}

func (h *Home) POST(Session *Server.Session) {

}

func (h *Home) DELETE(Session *Server.Session) {

}
