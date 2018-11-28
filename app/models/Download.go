package models

type Download struct {
	ID     	    		string `json:"id"`

	AllDebridID 		int `json:"allDebridId"`
	TorrentUrl		    string `json:"torrentUrl"`
	TorrentState		int `json:"torrentState"`
	DownloadState		int `json:"downloadState"`
	Percentage			float32 `json:"percentage"`
	Size				int `json:"size"`
	Speed				int `json:"speed"`

	Name				string `json:"name"`
	Type				string `json:"type"`
	Season				string `json:"season"`
	Episode				string `json:"episode"`
	Links				[]Link `json:"links"`
}

type Link struct {
	Url		    string `json:"url"`
	State		int `json:"state"`
	Percentage	float32 `json:"percentage"`
	Size		int `json:"size"`
	Speed		int `json:"speed"`
}

func GetLinksFromString(links []string, isDebrided bool) []Link {
	res := []Link{}
	for _, link := range links {
		var state int
		if (isDebrided) {
			state = LINK_QUEUING
		} else {
			state = LINK_NOT_DEBRIDED
		}

		res = append(res, Link{
			Percentage: 0,
			Size: 0,
			Speed: 0,
			State: state,
			Url: link,
		})
	}

	return res
}

const (
	TORRENT_QUEUING int = 0
	TORRENT_DOWNLOADING int = 1
	TORRENT_UPLOADING int = 2
	TORRENT_DONE int = 3

	DOWNLOAD_NOT_READY int = 0
	DOWNLOAD_DEBRIDING int = 1
	DOWNLOAD_QUEUING int = 2
	DOWNLOAD_DOWNLOADING int = 3
	DOWNLOAD_DECOMPRESSING int = 4
	DOWNLOAD_DONE int = 5

	LINK_NOT_DEBRIDED int = 0
	LINK_QUEUING int = 1
	LINK_DOWNLOADING int = 2
	LINK_DONE int = 3
)
