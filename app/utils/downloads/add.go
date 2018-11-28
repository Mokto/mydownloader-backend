package downloads

import (
	"localserver/app/models"
	"localserver/app/services/cache"
	"encoding/json"

)

func Add(download models.Download) error {

	cacheData, err := json.Marshal(download)
	if (err != nil) {
		return err
	}
	return cache.HSet(cacheKey, download.ID, cacheData)
}

// func Save(downloads []models.Download) error {
// 	for _, download := range downloads {
// 		cacheData, err := json.Marshal(download)
// 		if (err != nil) {
// 			return err
// 		}
// 		cache.HSet(cacheKey, download.ID, cacheData)
// 	}

// 	return nil
// }

