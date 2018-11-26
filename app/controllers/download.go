package controllers

import (
	"localserver/app/models"
	"fmt"
	"localserver/app/utils/debrid"
	"localserver/app/utils/links"

	"github.com/revel/revel"
	"github.com/satori/go.uuid"
)

// Download Controller
type Download struct {
	*revel.Controller
}

// Login to Download Provider
func (c Download) Download() revel.Result {

	var params map[string]string
	c.Params.BindJSON(&params)
	url := params["url"]
	name := params["name"]
	contentType := params["type"]
	episode := params["episode"]
	season := params["season"]
	c.Validation.Required(url)

	if c.Validation.HasErrors() {
		c.Response.Status = 400
		return c.RenderText("INVALID_PARAMS")
	}

	var debridInstance debrid.Debrid
	debridInstance = &debrid.AllDebrid{}

	err, allDebridID := debridInstance.AddTorrent("", url)
	if (err != nil) {
		fmt.Println(err)
		c.Response.Status = 500
		return c.RenderText(err.Error())
	}

	link := models.Link{
		ID:          		uuid.Must(uuid.NewV4()).String(),
		AllDebridID:		allDebridID,
		TorrentUrl:			url,
		Name:				name,
		Type:				contentType,
		Season:				season,
		Episode:			episode,
		TorrentState: 		models.TORRENT_QUEUING,
		DownloadState: 		models.DOWNLOAD_NOT_READY,
	}

	links.Add(link)

	go links.ListAndSend()
	
	return c.RenderJSON(nil)
}

// Login to Download Provider
func (c Download) DeleteDownload(id string) revel.Result {

	links.Remove(id)
	// TODO SHOULD REMOVE ALLDEBRID TOO

	go links.ListAndSend()
	
	return c.RenderJSON(nil)
}
