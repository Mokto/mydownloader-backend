package models

type Link struct {
	ID     	    		int `json:"id"`
	Url		    		string `json:"url"`
	Name				string `json:"name"`
	Type				string `json:"type"`
	Season				string `json:"season"`
	Episode				string `json:"episode"`
	TorrentDebriding	bool `json:"torrentDebriding"`
	LinkDownloading		bool `json:"linkDownloading"`
	Percentage			float32 `json:"percentage"`
	Size				int `json:"size"`
	Speed		 		int `json:"speed"`
}
