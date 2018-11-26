
package alldebrid

import (
	"encoding/json"
	"net/http"
	"strings"
)

type allDebridTorrentsGetResponse struct {
	Success    bool          			 `json:"success"`
	Torrents   []allDebridTorrent    `json:"torrents"`
}
type allDebridTorrent struct {
	ID   		int    	`json:"id"`
	Filename    string   	`json:"filename"`
	Size		int    		`json:"size"`
	StatusCode  int   		`json:"statusCode"`
	Downloaded  int  	    `json:"downloaded"`
	Uploaded   	int  		`json:"uploaded"`
	DownloadSpeed int   	`json:"downloadSpeed"`
	UploadSpeed  int    	`json:"uploadSpeed"`
	UploadDate 	int			`json:"uploadDate"`
	Links		[]string    `json:"link"`
}
func getTorrents() []allDebridTorrent {
	var url strings.Builder
	url.WriteString("https://api.alldebrid.com/user/torrents?agent=mySoft&token=")
	url.WriteString(getToken())

	httpResponse, _ := http.Get(url.String())
	
	allDebridTorrentsGetResponseR := allDebridTorrentsGetResponse{}
	json.NewDecoder(httpResponse.Body).Decode(&allDebridTorrentsGetResponseR)

	return allDebridTorrentsGetResponseR.Torrents
}


