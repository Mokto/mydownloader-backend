package controllers

import (
	"fmt"
	"localserver/app/utils/debrid"

	"github.com/revel/revel"
)

// Download Controller
type Download struct {
	*revel.Controller
}

// Login to Download Provider
func (c Download) Download() revel.Result {

	var params map[string]string
	c.Params.BindJSON(&params)
	link := params["link"]
	c.Validation.Required(link)

	if c.Validation.HasErrors() {
		c.Response.Status = 400
		return c.RenderText("INVALID_PARAMS")
	}

	fmt.Println(link)

	var debridInstance debrid.Debrid
	debridInstance = &debrid.AllDebrid{}

	// link, err := debridInstance.
	// fmt.Println(link)
	fmt.Println(err)
	// if err != nil {
	// 	fmt.Printf("All debrid HTTP error: %s\n", err)
	// 	c.Response.Status = 403
	// 	return c.RenderText("LOGIN_FAILED")
	// }

	return c.RenderJSON(nil)
}
