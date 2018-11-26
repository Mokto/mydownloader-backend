package links

import (
	"localserver/app/services/cache"
	"localserver/app/services/websockets"

)

func ListAndSend() {
	cacheData, _ := cache.Get(cacheKey).Result()
	if (cacheData == "") {
		cacheData = "[]"
	}
	websockets.SendMessage("links", cacheData)
}

