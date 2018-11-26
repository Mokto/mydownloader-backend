package runners

import (
	"fmt"
	"localserver/app/models"
	"localserver/app/utils/links"
	"localserver/app/utils/debrid"
	"time"
)


func WatchAllDebrid() {
	go func() {
		for {
			links := links.GetAll()
			for _, link := range links {
				if (link.AllDebridID != 0 && (link.TorrentState != models.TORRENT_DONE)) {
					fmt.Println("Updating statuses")
					debrid.UpdateStatuses(links)
					break
				}
			}
			time.Sleep(time.Second);
		}
	}()
}

