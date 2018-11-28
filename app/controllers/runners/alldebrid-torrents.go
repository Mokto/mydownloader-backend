package runners

import (
	"localserver/app/models"
	"localserver/app/utils/downloads"
	"localserver/app/utils/debrid"
	"time"
)


func WatchAllDebrid() {
	go func() {
		for {
			downloadsToCheck := downloads.GetAll()
			downloadsToUpdate := []models.Download{}
			for _, download := range downloadsToCheck {
				if (download.AllDebridID != 0 && (download.TorrentState != models.TORRENT_DONE)) {
					downloadsToUpdate = append(downloadsToUpdate, download)
				}
			}
			if (len(downloadsToUpdate) != 0) {
				debrid.UpdateStatuses(downloadsToUpdate)
			}
			time.Sleep(time.Second);
		}
	}()
}

