package controllers

import (
	"github.com/revel/revel"
)

// Torrents Controller
type Torrents struct {
	*revel.Controller
}

// Post receives torrent magnet and download it
func (c Torrents) Post() revel.Result {

	// var debridInstance debrid.Debrid
	// debridInstance = &debrid.AllDebrid{}

	// body, err := debridInstance.Login("moktoo", "medalist-vaporing-pillag")
	// if err != nil {
	// 	fmt.Printf("All debrid HTTP error: %s\n", err)
	// 	return c.RenderError(errors.New("AllDebrid login failed"))
	// }

	// fmt.Printf("%s", body)

	return c.RenderJSON(nil)
}
