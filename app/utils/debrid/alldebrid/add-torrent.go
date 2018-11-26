package alldebrid

import (
	"net/http"

	 "net/url"
	 "fmt"
)


func AddTorrent(filename string, magnet string) (error, int) {

	previousTorrents := getTorrents()
	_, err := http.PostForm("https://upload.alldebrid.com/uploadtorrent.php", url.Values{"uid": {getUid()}, "magnet": {magnet}, "splitfile": {"1"}, "quick": {"1"}})
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
			return nil, torrent.ID
		}
	}


	return err, 0
}
