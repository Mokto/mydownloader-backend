package debrid

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"localserver/app/services/cache"

	curl "github.com/andelf/go-curl"
	"github.com/usineur/go-debrid/utils"
)

var allDebridKey = "alldebrid"

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

	go cache.Set(allDebridKey, allDebridSuccessR.Token, 0)

	return
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
	cache.Delete(allDebridKey)
}

type getRedirectLinkResponse struct {
	Success bool `json:"success"`
	Folder  bool `json:"folder"`
	Links   bool `json:"links"`
}

// GetRedirectLink gets the final Link
// func (d *AllDebrid) GetRedirectLink(link string) (finalLink string, err error) {

// 	var url strings.Builder
// 	url.WriteString("https://api.alldebrid.com/link/redirector?agent=mySoft&token=")
// 	url.WriteString(getToken())
// 	url.WriteString("&link=")
// 	url.WriteString(link)

// 	fmt.Println(url.String())

// 	httpResponse, err := http.Get(url.String())

// 	if err != nil {
// 		return
// 	}
// 	defer httpResponse.Body.Close()

// 	bodyBytes, _ := ioutil.ReadAll(httpResponse.Body)
// 	bodyString := string(bodyBytes)
// 	fmt.Println(bodyString)

// 	getRedirectLinkResponseR := getRedirectLinkResponse{}
// 	json.NewDecoder(httpResponse.Body).Decode(&getRedirectLinkResponseR)

// 	fmt.Println(getRedirectLinkResponseR)

// 	// go cache.Set(allDebridKey, allDebridSuccessR.Token, 0)

// 	return
// }

func AddTorrent(filename string, magnet string) error {
	if uid, err := getUid(); err != nil {
		return err
	} else {
		path := "https://alldebrid.fr/torrent/"

		form := curl.NewForm()

		form.Add("uid", getToken())
		form.Add("domain", path)
		form.Add("magnet", magnet)
		if filename != "" {
			form.AddFile("files[]", filename)
		}
		form.Add("splitfile", utils.Btos(true))
		form.Add("quick", utils.Btos(true))
		form.Add("submit", "Convert this torrent")

		if res, eff, err := sendRequest("/uploadtorrent.php", nil, form); err != nil {
			return err
		} else if res == "Bad uploaded files" {
			return fmt.Errorf(res)
		} else if res == "Invalid cookie." {
			return fmt.Errorf(res)
		} else if pattern, err := regexp.Compile(host + path + "\\?error=(.*)"); err != nil {
			return err
		} else if matches := pattern.FindStringSubmatch(eff); len(matches) == 2 {
			switch matches[1] {
			case "":
				return fmt.Errorf("Alldebrid internal problem: retry")

			default:
				if msg, err := utils.GetContent(res, "//div[@style=\"color:red;text-align:center;\"]"); err != nil {
					return err
				} else {
					return fmt.Errorf(msg)
				}
			}
		} else if filename != "" {
			fmt.Printf("%v correctly added to torrent queue\n", filename)
			return nil
		} else {
			fmt.Println("magnet correctly added to torrent queue")
			return nil
		}
	}
}

func getToken() string {
	return cache.Get(allDebridKey).Val()
}
