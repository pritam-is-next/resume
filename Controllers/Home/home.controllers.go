package Home

import (
	components "github.com/pritam-is-next/resume/Components"
	Controller "github.com/vrianta/Server/Controller"
	"github.com/vrianta/Server/Session"
	"github.com/vrianta/Server/Template"
)

var Home = Controller.Struct{
	View: "home.php",
	GET:  GET,
}

var GET = func(self *Controller.Struct) *Template.Response {
	response := &Template.Response{
		"Title":          "Pritam Dutta",
		"Heading":        "Pritam Dutta",
		"NavItems":       components.NavItems,
		"Hero":           components.Hero,
		"AboutMe":        components.AboutMe,
		"Skills":         components.Skills,
		"Experiences":    components.Experiences,
		"Projects":       components.Projects,
		"ContactDetails": components.ContactDetails,
	}
	// if err := Session.RenderEngine.RenderTemplate("home.php", response); err != nil {
	// 	Session.RenderEngine.Render(err.Error())
	// }
	return response
}

func POST(Session *Session.Struct) {

}

func DELETE(Session *Session.Struct) {

}
