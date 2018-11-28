package alldebrid

import (
	"localserver/app/models"

	"localserver/app/utils/downloads"
)


func UpdateStatuses(downloadsToCheck []models.Download) {
	torrents := getTorrents()
	for  _, download := range downloadsToCheck {
		for  _, torrent := range torrents {
			if (download.AllDebridID == torrent.ID) {
				if (download.Name == "") {
					download.Name = torrent.Filename
				}
				download.Size = torrent.Size

				switch torrent.StatusCode {
					case 0:
						download.TorrentState = models.TORRENT_QUEUING
						download.Speed = 0
						download.Percentage = 0
					case 1:
						download.TorrentState = models.TORRENT_DOWNLOADING
						download.Speed = torrent.DownloadSpeed
						if (torrent.Size == 0) {
							download.Percentage = 0
						} else {
							download.Percentage = float32(torrent.Downloaded) * 100.0 / float32(torrent.Size)
						}
						
					case 2:
						download.TorrentState = models.TORRENT_DOWNLOADING
						download.Speed = 0
						download.Percentage = 100
					case 3:
						download.TorrentState = models.TORRENT_UPLOADING
						download.Speed = torrent.UploadSpeed
						download.Percentage = float32(torrent.Uploaded) * 100.0 / float32(torrent.Size)
					case 4:
						download.TorrentState = models.TORRENT_DONE
						download.DownloadState = models.DOWNLOAD_NOT_READY
						download.Speed = 0
						download.Percentage = 100
						download.Links = models.GetLinksFromString(torrent.Links, false)

						
				}

				downloads.Save(download)
			}
		}
	}

	downloads.ListAndSend()
}




