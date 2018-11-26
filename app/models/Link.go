package models

type Link struct {
	ID     	    		string `json:"id"`
	AllDebridID 		int `json:"allDebridId"`
	Url		    		string `json:"url"`
	Name				string `json:"name"`
	Type				string `json:"type"`
	Season				string `json:"season"`
	Episode				string `json:"episode"`
	TorrentDownloading	bool `json:"torrentDownloading"`
	TorrentUploading	bool `json:"torrentUploading"`
	LinkDownloading		bool `json:"linkDownloading"`
	Percentage			float32 `json:"percentage"`
	Size				int `json:"size"`
	Speed		 		int `json:"speed"`
}
