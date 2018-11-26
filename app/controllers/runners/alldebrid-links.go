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
			for  i, download := range downloadsToCheck {
				if download.DownloadState == models.DOWNLOAD_NOT_READY && download.TorrentState == models.TORRENT_DONE {
					downloadsToCheck[i].DownloadState = models.DOWNLOAD_DEBRIDING
					downloads.Save(downloadsToCheck)
					for  j, textdownload := range download.Links {
						downloadLink := debrid.GetDownloadableLink(textdownload)
						
						downloadsToCheck[i].Links[j] = downloadLink
					}
					downloadsToCheck[i].DownloadState = models.DOWNLOAD_DOWNLOADING
				}
			}
			downloads.Save(downloadsToCheck)
			downloads.ListAndSend()
			time.Sleep(time.Second);
		}
	}()
}
