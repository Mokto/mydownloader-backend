package debrid

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"



	"localserver/app/services/cache"
	 "net/url"
	 "fmt"
	"gopkg.in/headzoo/surf.v1"
	"gopkg.in/headzoo/surf.v1/agent"
	// "gopkg.in/headzoo/surf.v1/jar"
	// "gopkg.in/headzoo/surf.v1/browser"
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


func (d *AllDebrid) AddTorrent(filename string, magnet string) error {

	_, err := http.PostForm("https://upload.alldebrid.com/uploadtorrent.php", url.Values{"uid": {getUid()}, "magnet": {magnet}, "splitfile": {"1"}, "quick": {"1"}})

	return err
}

func getToken() string {
	return cache.Get(allDebridTokenKey).Val()
}
func getUid() string {
	return cache.Get(allDebridUidKey).Val()
}


// // Login to AllDebrid to be able to make requests
// func (d *AllDebrid) GetTorrentsStatus() (err error) {

// 	var url strings.Builder
// 	url.WriteString("https://api.alldebrid.com/user/login?agent=mySoft&username=")
// 	url.WriteString(username)
// 	url.WriteString("&password=")
// 	url.WriteString(password)

// 	httpResponse, err := http.Get(url.String())

// 	if err != nil {
// 		return
// 	}
// 	defer httpResponse.Body.Close()

// 	if httpResponse.StatusCode != 200 {
// 		if httpResponse.StatusCode == 429 {
// 			err = errors.New("Too many requests")
// 			return
// 		}

// 		allDebridErrorR := allDebridLoginError{}
// 		json.NewDecoder(httpResponse.Body).Decode(&allDebridErrorR)

// 		err = errors.New(allDebridErrorR.Error)
// 		return
// 	}

// 	allDebridSuccessR := allDebridLoginSuccess{}
// 	json.NewDecoder(httpResponse.Body).Decode(&allDebridSuccessR)

// 	go cache.Set(allDebridTokenKey, allDebridSuccessR.Token, 0)
// 	go webLogin(username, password)

// 	return
// }