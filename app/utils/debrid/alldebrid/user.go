package alldebrid

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"



	"localserver/app/services/cache"
	 "net/url"

	"gopkg.in/headzoo/surf.v1"
	"gopkg.in/headzoo/surf.v1/agent"
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

// Login to AllDebrid to be able to make requests
func Login(username, password string) (err error) {

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
			go cache.Set(allDebridUidKey, cookie.Value, 0)
		}
		
	}

}

// IsLoggedIn returns a boolean which says if you are logged in or not
func IsLoggedIn() bool {
	if getToken() != "" {
		return true
	}

	return false
}

// Logout removes the Alldebrid cache token
func Logout() {
	cache.Delete(allDebridTokenKey)
}

func getToken() string {
	return cache.Get(allDebridTokenKey).Val()
}
func getUid() string {
	return cache.Get(allDebridUidKey).Val()
}


