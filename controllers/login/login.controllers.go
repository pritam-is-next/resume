package login

import (
	components "github.com/pritam-is-next/resume/components"
	models "github.com/pritam-is-next/resume/models"
	"github.com/vrianta/agai/v1/controller"
	"github.com/vrianta/agai/v1/log"
	"github.com/vrianta/agai/v1/utils"
)

type Controller struct {
	controller.Context
}

func (c *Controller) GET() controller.View {
	if c.IsLoggedIn() {
		log.Debug("User is Alreadt Logged In in Login GET Method")
		c.Redirect("/admin")
		return controller.EmptyResponse().ToView("login")
	}

	response := controller.Response{
		"Title":          "Pritam Dutta",
		"Heading":        "Pritam Dutta",
		"NavItems":       models.Nav_items.GetComponents(),
		"Hero":           components.Hero,
		"AboutMe":        components.AboutMe,
		"Skills":         components.Skills,
		"Experiences":    components.Experiences,
		"Projects":       components.Projects,
		"ContactDetails": components.ContactDetails,
	}
	return response.ToView("login")
}

func (c *Controller) POST() controller.View {
	if c.IsLoggedIn() {
		c.Redirect("/admin")
		log.WriteLog("Redirecting to Admin")
		return controller.EmptyResponse().ToView("login")
	}

	email := c.GetInput("loginEmail")
	password := c.GetInput("loginPassword")

	if email == nil || password == nil {
		r := controller.Response{
			"error": "email or password field is empty",
		}
		return r.ToView("login")
	}

	if user, err := models.Users.Get().
		Where(models.Users.Fields.UserName).Is(email.(string)).
		First(); err != nil {
		log.WriteLog("Got error while fetching ", err.Error())
		r := controller.Response{
			"email":    email,
			"password": password,
			"error":    err.Error(),
		}
		return r.ToView("login")
	} else if user != nil && utils.CheckPassword(user["Password"].(string), password.(string)) {
		log.Debug("Successfully Logged in")
		c.Login()
		c.Redirect("/")
	} else {
		r := controller.Response{
			"UserName": email,
			"Password": password,
			"error":    "User Name or Password is wrong",
		}
		return r.ToView("login")
	}

	return controller.EmptyResponse().ToView("")
}
