package downloads

import (
	// "fmt"
	"encoding/json"
	"localserver/app/services/cache"
	"localserver/app/models"
	"localserver/app/services/websockets"

)

func ListAndSend() {
	downloads := GetAll()
	send(downloads)
}

func send(downloads []models.Download) {
	cacheData, _ := json.Marshal(downloads)
	websockets.SendMessage("downloads", string(cacheData))
}


func GetAll() (downloads []models.Download) {
	downloads = []models.Download{}
	cacheData, err := cache.HGetAll(cacheKey)
	if (cacheData == nil || err != nil) {
		return;
	}
	for _, value := range cacheData {
		download := models.Download{}
		json.Unmarshal([]byte(value), &download) 
        downloads = append(downloads, download)
    }

	return;
}


func Get(id string) (download models.Download) {

	cacheData, err := cache.HGet(cacheKey, id)
	if (cacheData == "" || err == nil) {
		return;
	}
	json.Unmarshal([]byte(cacheData), &download) 
	return
}