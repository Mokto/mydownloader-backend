package debrid

import (
	"localserver/app/models"
	"encoding/json"
	"errors"
	"net/http"
	"strings"



	"localserver/app/services/cache"
	"localserver/app/utils/links"
	 "net/url"
	 "fmt"
	"gopkg.in/headzoo/surf.v1"
	"gopkg.in/headzoo/surf.v1/agent"
	"io/ioutil"
)

var allDebridTokenKey = "alldebrid_token"
var allDebridUidKey = "alldebrid_uid"

type userAllDebrid struct {
	username     string
	email        string
	isPremium    bool
	premiumUntil int
}

type allDebridLoginSuccess struct {
	Success bool          `json:"success"`
	Token   string        `json:"token"`
	User    userAllDebrid `json:"user"`
}

type allDebridLoginError struct {
	Error     string `json:"error"`
	ErrorCode int    `json:"errorCode"`
}

// AllDebrid Service
type AllDebrid struct{}

// Login to AllDebrid to be able to make requests
func (d *AllDebrid) Login(username, password string) (err error) {

	var url strings.Builder
	url.WriteString("https://api.alldebrid.com/user/login?agent=mySoft&username=")
	url.WriteString(username)
	url.WriteString("&password=")
	url.WriteString(password)

	httpResponse, err := http.Get(url.String())

	if err != nil {
		return
	}
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != 200 {
		if httpResponse.StatusCode == 429 {
			err = errors.New("Too many requests")
			return
		}

		allDebridErrorR := allDebridLoginError{}
		json.NewDecoder(httpResponse.Body).Decode(&allDebridErrorR)

		err = errors.New(allDebridErrorR.Error)
		return
	}

	allDebridSuccessR := allDebridLoginSuccess{}
	json.NewDecoder(httpResponse.Body).Decode(&allDebridSuccessR)

	go cache.Set(allDebridTokenKey, allDebridSuccessR.Token, 0)
	go webLogin(username, password)

	return
}


func webLogin(username, password string) {
	bow := surf.NewBrowser()
	bow.SetUserAgent(agent.Chrome())
	err := bow.Open("https://alldebrid.fr/register/")
    if err != nil {
		return
	}
	fm, _ := bow.Form("form#loginForm")
    fm.Input("login_login", username)
    fm.Input("login_password", password)
    if fm.Submit() != nil {
        return
	}
	url, _ := url.Parse("https://alldebrid.com")
	for _, cookie := range bow.CookieJar().Cookies(url) {
		if (cookie.Name == "uid") {
			fmt.Println(allDebridUidKey)
			fmt.Println(cookie.Value)
			go cache.Set(allDebridUidKey, cookie.Value, 0)
		}
		
	}

}

// IsLoggedIn returns a boolean which says if you are logged in or not
func (d *AllDebrid) IsLoggedIn() bool {
	if getToken() != "" {
		return true
	}

	return false
}

// Logout removes the Alldebrid cache token
func (d *AllDebrid) Logout() {
	cache.Delete(allDebridTokenKey)
}


func (d *AllDebrid) AddTorrent(filename string, magnet string) (error, int) {

	previousTorrents := getTorrents()
	resp, err := http.PostForm("https://upload.alldebrid.com/uploadtorrent.php", url.Values{"uid": {getUid()}, "magnet": {magnet}, "splitfile": {"1"}, "quick": {"1"}})
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(bodyBytes))
	if (err != nil) {
		return err, 0
	}
	newTorrents := getTorrents()

	if (len(previousTorrents) == len(newTorrents)) {
		if (len(newTorrents) == 0) {
			return fmt.Errorf("ERROR_ADDING_TORRENT"), 0	
		}
		return fmt.Errorf("ALREADY_EXISTING"), 0
	}

	// Get the torrent which was not present before
	torrentsMap := map[int]allDebridTorrent{}
	for _, torrent := range previousTorrents {
		torrentsMap[torrent.ID] = torrent;
	}
	for _, torrent := range newTorrents {
		if _, ok := torrentsMap[torrent.ID]; !ok {
			fmt.Println(torrent)
			return nil, torrent.ID
		}
	}


	return err, 0
}

func (d *AllDebrid) UpdateStatuses(linksToCheck []models.Link) {
	if (linksToCheck == nil) {
		linksToCheck = links.GetAll()
	}
	torrents := getTorrents()
	// fmt.Println(torrents)
	for  _, torrent := range torrents {
		for  i, link := range linksToCheck {
			if (link.AllDebridID == torrent.ID) {
				if (link.Name == "") {
					linksToCheck[i].Name = torrent.Filename
				}
				linksToCheck[i].Size = torrent.Size
				linksToCheck[i].TorrentDownloading = torrent.Size == 0 || torrent.Downloaded != torrent.Size
				linksToCheck[i].TorrentUploading = !linksToCheck[i].TorrentDownloading && torrent.Uploaded != torrent.Size

				if (linksToCheck[i].TorrentDownloading) {
					linksToCheck[i].Speed = torrent.DownloadSpeed
				} else if (linksToCheck[i].TorrentUploading) {
					linksToCheck[i].Speed = torrent.UploadSpeed
				}

				if (linksToCheck[i].TorrentDownloading) {
					if (torrent.Size == 0) {
						linksToCheck[i].Percentage = 0
					} else {
						var percentage float32 = float32(torrent.Downloaded) * 100.0 / float32(torrent.Size);
						fmt.Println(percentage)
						linksToCheck[i].Percentage = float32(torrent.Downloaded) * 100.0 / float32(torrent.Size)
					}
					
				}
				if (linksToCheck[i].TorrentUploading) {
					linksToCheck[i].Percentage = float32(torrent.Uploaded) * 100.0 / float32(torrent.Size)
				}
				fmt.Println(torrent.Uploaded, torrent.Size, linksToCheck[i].Name, linksToCheck[i].TorrentDownloading, linksToCheck[i].TorrentUploading, linksToCheck[i].Percentage )
			}
		}
	}

	links.Save(linksToCheck)
	links.Send(linksToCheck)
}

func getToken() string {
	return cache.Get(allDebridTokenKey).Val()
}
func getUid() string {
	return cache.Get(allDebridUidKey).Val()
}


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
