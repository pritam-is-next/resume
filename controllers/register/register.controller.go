package register

import (
	"github.com/pritam-is-next/resume/models"
	"github.com/vrianta/agai/v1/controller"
	"github.com/vrianta/agai/v1/log"
	"github.com/vrianta/agai/v1/utils"
)

/*
Register Controller

This is a basic controller built on top of `controller.Context`. It supports all the usual HTTP methods:
GET, POST, PUT, DELETE, PATCH, HEAD, and OPTIONS.

Each method is just a variable on the struct — you can set any of them to a handler function like:

	func(self *controller.Context) *template.Response

If you don't define a method, GET will be used by default. So if your controller just needs to respond to GET requests,
you only need to set that and you're good to go.

This setup keeps things simple — define what you need, skip what you don't.
*/

type Controller struct {
	controller.Context
}

func (c *Controller) GET() controller.View {

	component, ok := models.App_state.GetComponent("initialised")

	if !ok {
		panic("Needed Component Missing")
	}

	val, ok := component.FieldValue("Value")
	if !ok {
		panic("Needed Component Missing")
	}

	initialised, ok := val.(string)

	if !ok {
		panic("Needed Component Missing")
	}

	log.Info("%s", initialised)
	if initialised == "true" {
		c.Redirect("/")
	}

	return controller.EmptyResponse().ToView("register")
}

func (c *Controller) POST() controller.View {
	initialised, i_ok := models.App_state.GetComponent("initialised")
	if !i_ok {
		c.WithCode("/register", controller.HttpStatus.InternalServerError)
	}

	first_name, first_name_ok := c.GetInput("firstName").(string)
	last_name, last_name_ok := c.GetInput("lastName").(string)
	email, email_ok := c.GetInput("email").(string)
	password, password_ok := c.GetInput("password").(string)
	confirmPassword, confirmPassword_ok := c.GetInput("confirmPassword").(string)

	if !first_name_ok || !last_name_ok || !email_ok || !password_ok || !confirmPassword_ok {
		// not ok
		return (&controller.Response{
			"error": "Please fill the values correctly",
		}).ToView("register")
	}

	if password != confirmPassword {
		return (&controller.Response{
			"error": "Password and Confirm password do not match",
		}).ToView("register")
	}

	if hashed_password, err := utils.HashPassword(password); err != nil {
		return (&controller.Response{
			"error": "Internal server error | failed to hash password",
		}).ToView("register")

	} else if err := models.Users.Create().
		Set(models.Users.Fields.UserId).To(email).
		Set(models.Users.Fields.UserName).To(email).
		Set(models.Users.Fields.Password).To(hashed_password).
		Set(models.Users.Fields.FirstName).To(first_name).
		Set(models.Users.Fields.LastName).To(last_name).Exec(); err != nil {
		return (&controller.Response{
			"error": "Internal server error | failed to Create User Table " + err.Error(),
		}).ToView("register")
	}

	if err := models.User_details.Create().
		Set(models.User_details.Fields.UserId).To(email).
		Set(models.User_details.Fields.FullName).To(first_name + " " + last_name).
		Set(models.User_details.Fields.AboutMe).To("I am Human").Exec(); err != nil {
		return (&controller.Response{
			"error": "Internal server error | failed to Create User Details Table " + err.Error(),
		}).ToView("register")
	}

	initialised["Value"] = "t"

	models.App_state.UpdateComponent("initialised", initialised)

	// models.App_state
	c.Redirect("/")

	return controller.EmptyResponse().ToView("register")

}
