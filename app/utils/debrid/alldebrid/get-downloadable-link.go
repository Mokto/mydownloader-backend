package alldebrid

import (
	"encoding/json"
	"net/http"
	"strings"
)



type allDebridGetDownloadableLinkInfoResponse struct {
	Link string `json:"link"`
}
type allDebridGetDownloadableLinkResponse struct {
	Success bool `json:"success"`
	Infos allDebridGetDownloadableLinkInfoResponse `json:"infos"`
}

func GetDownloadableLink(link string) string {
	var url strings.Builder
	url.WriteString("https://api.alldebrid.com/link/unlock?agent=mySoft&token=")
	url.WriteString(getToken())
	url.WriteString("&link=")
	url.WriteString(link)

	httpResponse, _ := http.Get(url.String())
	
	allDebridGetDownloadableLinkResponseR := allDebridGetDownloadableLinkResponse{}
	json.NewDecoder(httpResponse.Body).Decode(&allDebridGetDownloadableLinkResponseR)

	return allDebridGetDownloadableLinkResponseR.Infos.Link
}
