package runners

import (
	"fmt"
	"localserver/app/models"
	"localserver/app/utils/downloads"
	"localserver/app/utils/downloadmanager"
	"time"
)


func Download() {
	go func() {
		for {
			downloadsToCheck := downloads.GetAll()
			for i, download := range downloadsToCheck {
				if (download.DownloadState == models.DOWNLOAD_QUEUING) {
					fmt.Println("Download the file")
					downloadmanager.Download(download)
					downloadsToCheck[i].DownloadState = models.DOWNLOAD_DOWNLOADING
					downloadsToCheck[i].Percentage = 0
					downloads.Save(downloadsToCheck)
					break
				}
			}
			time.Sleep(time.Second);
		}
	}()
}

