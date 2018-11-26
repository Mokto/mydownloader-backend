package debrid

import "localserver/app/models"
import "localserver/app/utils/debrid/alldebrid"

// 	UpdateStatuses(links []models.Link)
// 	GetDownloadableLink(link string) string

func Login(username, password string) error {
	return alldebrid.Login(username, password)
}

func Logout() {
	alldebrid.Logout()
}

func AddTorrent(filename string, magnet string) (error, int) {
	return alldebrid.AddTorrent(filename, magnet)
}

func IsLoggedIn() bool {
	return alldebrid.IsLoggedIn()
}

func UpdateStatuses(links []models.Link) {
	alldebrid.UpdateStatuses(links)
}

func GetDownloadableLink(link string) string {
	return alldebrid.GetDownloadableLink(link)
}