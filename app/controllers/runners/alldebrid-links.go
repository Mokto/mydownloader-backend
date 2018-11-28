package runners

import (
	"localserver/app/models"
	"localserver/app/utils/downloads"
	"localserver/app/utils/debrid"
	"time"
)

func CheckLinksToDebrid() {
	go func() {
		for {
			downloadsToCheck := downloads.GetAll()
			hasChanged := false
			for  _, download := range downloadsToCheck {
				if download.DownloadState == models.DOWNLOAD_NOT_READY && download.TorrentState == models.TORRENT_DONE {
					download.DownloadState = models.DOWNLOAD_DEBRIDING
					downloads.Save(download)
					for  j, textdownload := range download.Links {
						downloadLink := debrid.GetDownloadableLink(textdownload.Url)
						
						download.Links[j].Url = downloadLink
						download.Links[j].State = models.LINK_QUEUING
					}
					download.DownloadState = models.DOWNLOAD_QUEUING
					downloads.Save(download)
					hasChanged = true
				}
			}
			if hasChanged == true {
				downloads.ListAndSend()
			}
			time.Sleep(time.Second);
		}
	}()
}
