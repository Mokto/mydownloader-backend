package runners

import (
	"fmt"
	"localserver/app/models"
	"localserver/app/utils/links"
	"localserver/app/utils/debrid"
	"time"
)

func CheckLinksToDebrid() {
	go func() {
		for {
			linksToDebrid := links.GetAll()
			for  i, link := range linksToDebrid {
				if link.DownloadState == models.DOWNLOAD_NOT_READY && link.TorrentState == models.TORRENT_DONE {
					linksToDebrid[i].DownloadState = models.DOWNLOAD_DEBRIDING
					links.Save(linksToDebrid)
					for  j, textLink := range link.Links {
						fmt.Println(textLink)
						linkDownloadable := debrid.GetDownloadableLink(textLink)
						
						linksToDebrid[i].Links[j] = linkDownloadable
					}
					linksToDebrid[i].DownloadState = models.DOWNLOAD_DOWNLOADING
				}
			}
			links.Save(linksToDebrid)
			links.ListAndSend()
			time.Sleep(time.Second);
		}
	}()
}
