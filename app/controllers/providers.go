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

	var debridInstance debrid.Debrid
	debridInstance = &debrid.AllDebrid{}

	data := make(map[string]bool)
	data["allDebrid"] = debridInstance.IsLoggedIn()

	return c.RenderJSON(data)
}
