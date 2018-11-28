package downloads

import (
	"localserver/app/models"
	"localserver/app/services/cache"
	"encoding/json"

)

func Save(download models.Download) error {
	cacheData, err := json.Marshal(download)
	if (err != nil) {
		return err
	}
	cache.HSet(cacheKey, download.ID, cacheData)

	return nil
}
