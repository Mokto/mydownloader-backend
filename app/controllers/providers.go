package controllers

import (
	"localserver/app/utils/debrid"

	"github.com/revel/revel"
)

// Providers Controller
type Providers struct {
	*revel.Controller
}

type providersStatus struct {
	allDebrid bool
}

// Status returns provider status
func (c Providers) Status() revel.Result {

	data := make(map[string]bool)
	data["allDebrid"] = debrid.IsLoggedIn()

	return c.RenderJSON(data)
}
