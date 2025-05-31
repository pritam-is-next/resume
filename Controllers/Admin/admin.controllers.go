// package Controllers

// import (
// 	components "github.com/pritam-is-next/resume/Components"
// 	Server "github.com/vrianta/Server"
// 	"github.com/vrianta/Server/Controller"
// )

// type Admin Controller.Struct

// func (h *Admin) Demo(Session *Server.Session) {

// 	if !Session.IsLoggedIn() {
// 		Session.Redirect("/login")
// 	}

// 	if err := Session.RenderEngine.RenderTemplate("admin.php", Server.TemplateData{
// 		"Title":          "Pritam Dutta",
// 		"Heading":        "Pritam Dutta",
// 		"NavItems":       components.NavItems,
// 		"Hero":           components.Hero,
// 		"AboutMe":        components.AboutMe,
// 		"Skills":         components.Skills,
// 		"Experiences":    components.Experiences,
// 		"Projects":       components.Projects,
// 		"ContactDetails": components.ContactDetails,
// 	}); err != nil {
// 		Session.RenderEngine.Render(err.Error())
// 	}
// }

// func (h *Admin) POST(Session *Server.Session) {

// }

// func (h *Admin) DELETE(Session *Server.Session) {

// }
