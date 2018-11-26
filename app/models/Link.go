package models

type Link struct {
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
}

const (
	TORRENT_QUEUING int = 0
	TORRENT_DOWNLOADING int = 1
	TORRENT_UPLOADING int = 2
	TORRENT_DONE int = 3

	DOWNLOAD_NOT_READY int = 0
	DOWNLOAD_QUEUING int = 1
	DOWNLOAD_DOWNLOADING int = 2
	DOWNLOAD_DONE int = 3
)
