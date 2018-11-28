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
			hasChanged := false
			for _, download := range downloadsToCheck {
				if (download.DownloadState == models.DOWNLOAD_QUEUING) {
					fmt.Println("Download the file")
					downloadmanager.Download(download)
					download.DownloadState = models.DOWNLOAD_DOWNLOADING
					download.Percentage = 0
					downloads.Save(download)
					hasChanged = true
				}
			}
			if (hasChanged) {
				downloads.ListAndSend()
			}
			time.Sleep(time.Second);
		}
	}()
}

