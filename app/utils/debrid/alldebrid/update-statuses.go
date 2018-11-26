package alldebrid

import (
	"localserver/app/models"

	"localserver/app/utils/links"
)


func UpdateStatuses(linksToCheck []models.Link) {
	if (linksToCheck == nil) {
		linksToCheck = links.GetAll()
	}
	torrents := getTorrents()
	for  _, torrent := range torrents {
		for  i, link := range linksToCheck {
			if (link.AllDebridID == torrent.ID) {
				if (link.Name == "") {
					linksToCheck[i].Name = torrent.Filename
				}
				linksToCheck[i].Size = torrent.Size

				switch torrent.StatusCode {
					case 0:
						linksToCheck[i].TorrentState = models.TORRENT_QUEUING
						linksToCheck[i].Speed = 0
						linksToCheck[i].Percentage = 0
					case 1:
						linksToCheck[i].TorrentState = models.TORRENT_DOWNLOADING
						linksToCheck[i].Speed = torrent.DownloadSpeed
						if (torrent.Size == 0) {
							linksToCheck[i].Percentage = 0
						} else {
							linksToCheck[i].Percentage = float32(torrent.Downloaded) * 100.0 / float32(torrent.Size)
						}
						
					case 2:
						linksToCheck[i].TorrentState = models.TORRENT_DOWNLOADING
						linksToCheck[i].Speed = 0
						linksToCheck[i].Percentage = 100
					case 3:
						linksToCheck[i].TorrentState = models.TORRENT_UPLOADING
						linksToCheck[i].Speed = torrent.UploadSpeed
						linksToCheck[i].Percentage = float32(torrent.Uploaded) * 100.0 / float32(torrent.Size)
					case 4:
						linksToCheck[i].TorrentState = models.TORRENT_DONE
						// linksToCheck[i].DownloadState = models.DOWNLOAD_NOT_DEBRIDED
						linksToCheck[i].Speed = 0
						linksToCheck[i].Percentage = 100
						linksToCheck[i].Links = torrent.Links
				}
			}
		}
	}

	links.Save(linksToCheck)
	links.Send(linksToCheck)

}




