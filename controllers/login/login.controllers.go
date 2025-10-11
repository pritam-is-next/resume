package login

import (
	components "github.com/pritam-is-next/resume/components"
	models "github.com/pritam-is-next/resume/models"
	"github.com/vrianta/agai/v1"
	"github.com/vrianta/agai/v1/log"
	"github.com/vrianta/agai/v1/utils"
)

type Controller struct {
	agai.Controller
}

func (c *Controller) GET() agai.View {
	if c.IsLoggedIn() {
		log.Debug("User is Alreadt Logged In in Login GET Method")
		c.Redirect("/admin")
		return c.View("login", c.EmptyResponse())
	}

	c.GetInputs()

	response := agai.Response{
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
	return c.View("login", response)
}

func (c *Controller) POST() agai.View {
	if c.IsLoggedIn() {
		log.WriteLog("Redirecting to Admin")
		return c.Redirect("/admin")
	}

	email := c.GetInput("loginEmail")
	password := c.GetInput("loginPassword")

	if email == nil || password == nil {
		r := agai.Response{
			"error": "email or password field is empty",
		}
		return c.View("login", r)
	}

	if user, err := models.Users.Get().
		Where(models.Users.Fields.UserName).Is(email.(string)).
		First(); err != nil {
		log.WriteLog("Got error while fetching ", err.Error())
		r := agai.Response{
			"email":    email,
			"password": password,
			"error":    err.Error(),
		}
		return c.View("login", r)
	} else if user != nil && utils.CheckPassword(user["Password"].(string), password.(string)) {
		log.Debug("Successfully Logged in")
		c.Login()
		c.Redirect("/home")
	} else {
		r := agai.Response{
			"UserName": email,
			"Password": password,
			"error":    "User Name or Password is wrong",
		}
		return c.View("login", r)
	}

	return c.View("home", c.EmptyResponse())
}
