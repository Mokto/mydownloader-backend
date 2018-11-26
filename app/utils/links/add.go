package links

import (
	"localserver/app/models"
	"localserver/app/services/cache"
	"encoding/json"

)

func Add(link models.Link) error {
	var links = GetAll()
	links = append(links, link)
	return Save(links)
}

func Save(links []models.Link) error {
	cacheBytes, err := json.Marshal(links)
	if (err != nil) {
		return err
	} 
	cache.Set(cacheKey, cacheBytes, 0)

	return nil
}

