package debrid

import "localserver/app/models"

// Debrid is the general type for every debrid provider
type Debrid interface {
	Login(username, password string) error
	IsLoggedIn() bool
	Logout()
	AddTorrent(filename string, magnet string) (error, int)
	UpdateStatuses(links []models.Link)
}
