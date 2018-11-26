package runners

import (
	"fmt"
	"localserver/app/models"
	"localserver/app/utils/downloads"
	"localserver/app/utils/debrid"
	"time"
)


func WatchAllDebrid() {
	go func() {
		for {
			downloadsToCheck := downloads.GetAll()
			for _, download := range downloadsToCheck {
				if (download.AllDebridID != 0 && (download.TorrentState != models.TORRENT_DONE)) {
					fmt.Println("Updating statuses")
					debrid.UpdateStatuses(downloadsToCheck)
					break
				}
			}
			time.Sleep(time.Second);
		}
	}()
}

