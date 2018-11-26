package downloads

import (
	"encoding/json"
	"localserver/app/services/cache"
	"localserver/app/models"
	"localserver/app/services/websockets"

)

func ListAndSend() {
	cacheData, _ := cache.Get(cacheKey).Result()
	if (cacheData == "") {
		cacheData = "[]"
	}
	websockets.SendMessage("downloads", cacheData)
}

func Send(downloads []models.Download) {
	cacheData, _ := json.Marshal(downloads)
	websockets.SendMessage("downloads", string(cacheData))
}


func GetAll() (downloads []models.Download) {
	downloads = []models.Download{}

	cacheData, _ := cache.Get(cacheKey).Result()
	if (cacheData != "") {
		json.Unmarshal([]byte(cacheData), &downloads) 
	}

	return;
}