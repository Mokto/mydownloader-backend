package alldebrid

import (
	"fmt"
	"strings"
	"time"
	"net/url"
	"net/http/cookiejar"
	"net/http"
	 "gopkg.in/headzoo/surf.v1"
	 "gopkg.in/headzoo/surf.v1/agent"
)


func RemoveTorrent(allDebridID int) (error) {
	var httpUrl strings.Builder
	httpUrl.WriteString("https://alldebrid.com//torrent/?action=remove&id=")
	httpUrl.WriteString(string(allDebridID))

	fmt.Println(httpUrl.String())
	fmt.Println(allDebridID)

	bow := surf.NewBrowser()
	bow.SetUserAgent(agent.Chrome())
	url, _ := url.Parse("https://alldebrid.com")
	cookieJar, _  := cookiejar.New(nil)
	uidCookie := http.Cookie{Name: "uid", Value: getUid() ,Expires: time.Now().Add(365*24*time.Hour)}
	cookies := []*http.Cookie{&uidCookie}
	cookieJar.SetCookies(url, cookies)
	bow.SetCookieJar(cookieJar)
	return bow.Open(httpUrl.String())
}
