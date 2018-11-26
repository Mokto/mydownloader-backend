package downloads

import (
	"localserver/app/models"
	"localserver/app/services/cache"
	"encoding/json"

)

func Add(download models.Download) error {
	var downloads = GetAll()
	downloads = append(downloads, download)
	return Save(downloads)
}

func Save(downloads []models.Download) error {
	cacheBytes, err := json.Marshal(downloads)
	if (err != nil) {
		return err
	} 
	cache.Set(cacheKey, cacheBytes, 0)

	return nil
}

