package links

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
	websockets.SendMessage("links", cacheData)
}

func Send(links []models.Link) {
	cacheData, _ := json.Marshal(links)
	websockets.SendMessage("links", string(cacheData))
}


func GetAll() (links []models.Link) {
	links = []models.Link{}

	cacheData, _ := cache.Get(cacheKey).Result()
	if (cacheData != "") {
		json.Unmarshal([]byte(cacheData), &links) 
	}

	return;
}