package links

import (
	"localserver/app/models"
	"localserver/app/services/cache"
	"encoding/json"

)

func Add(link models.Link) error {

	var links = []models.Link{}

	cacheData, _ := cache.Get(cacheKey).Result()
	if (cacheData != "") {
		json.Unmarshal([]byte(cacheData), &links) 
	}
	links = append(links, link)
	
	cacheBytes, err := json.Marshal(links)
	if (err != nil) {
		return err
	} 
	cache.Set(cacheKey, cacheBytes, 0)

	return nil
}

