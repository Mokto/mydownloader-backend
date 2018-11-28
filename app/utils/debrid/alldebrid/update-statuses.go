package alldebrid

import (
	"localserver/app/models"

	"localserver/app/utils/downloads"
)


func UpdateStatuses(downloadsToCheck []models.Download) {
	if (downloadsToCheck == nil) {
		downloadsToCheck = downloads.GetAll()
	}
	torrents := getTorrents()
	for  _, torrent := range torrents {
		for  i, link := range downloadsToCheck {
			if (link.AllDebridID == torrent.ID) {
				if (link.Name == "") {
					downloadsToCheck[i].Name = torrent.Filename
				}
				downloadsToCheck[i].Size = torrent.Size

				switch torrent.StatusCode {
					case 0:
						downloadsToCheck[i].TorrentState = models.TORRENT_QUEUING
						downloadsToCheck[i].Speed = 0
						downloadsToCheck[i].Percentage = 0
					case 1:
						downloadsToCheck[i].TorrentState = models.TORRENT_DOWNLOADING
						downloadsToCheck[i].Speed = torrent.DownloadSpeed
						if (torrent.Size == 0) {
							downloadsToCheck[i].Percentage = 0
						} else {
							downloadsToCheck[i].Percentage = float32(torrent.Downloaded) * 100.0 / float32(torrent.Size)
						}
						
					case 2:
						downloadsToCheck[i].TorrentState = models.TORRENT_DOWNLOADING
						downloadsToCheck[i].Speed = 0
						downloadsToCheck[i].Percentage = 100
					case 3:
						downloadsToCheck[i].TorrentState = models.TORRENT_UPLOADING
						downloadsToCheck[i].Speed = torrent.UploadSpeed
						downloadsToCheck[i].Percentage = float32(torrent.Uploaded) * 100.0 / float32(torrent.Size)
					case 4:
						downloadsToCheck[i].TorrentState = models.TORRENT_DONE
						downloadsToCheck[i].DownloadState = models.DOWNLOAD_NOT_READY
						downloadsToCheck[i].Speed = 0
						downloadsToCheck[i].Percentage = 100
						downloadsToCheck[i].Links = models.GetLinksFromString(torrent.Links, false)

						
				}
			}
		}
	}

	downloads.Save(downloadsToCheck)
	downloads.Send(downloadsToCheck)

}




