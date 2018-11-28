package downloads

import "localserver/app/services/cache"


func Remove(id string) error {
	return cache.HDel(cacheKey, id)
}
