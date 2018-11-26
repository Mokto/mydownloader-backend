package controllers

import (
	"fmt"
	"localserver/app/utils/debrid"

	"github.com/revel/revel"
)

// AllDebrid Controller
type AllDebrid struct {
	*revel.Controller
}

// Login to AllDebrid Provider
func (c AllDebrid) Login() revel.Result {

	var params map[string]string
	c.Params.BindJSON(&params)
	username := params["username"]
	password := params["password"]
	c.Validation.Required(username)
	c.Validation.Required(password)

	if c.Validation.HasErrors() {
		c.Response.Status = 400
		return c.RenderText("INVALID_PARAMS")
	}

	err := debrid.Login(username, password)
	if err != nil {
		fmt.Printf("All debrid HTTP error: %s\n", err)
		c.Response.Status = 403
		return c.RenderText("LOGIN_FAILED")
	}

	return c.RenderJSON(nil)
}

// Logout from AllDebrid
func (c AllDebrid) Logout() revel.Result {
	go debrid.Logout()

	return c.RenderJSON(nil)
}
