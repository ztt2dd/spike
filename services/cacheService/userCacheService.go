package cacheService

import (
	"spikeKill/pkg/e"
	"strconv"
)

func GetUserKey(id int) string {
	return e.CACHE_USER + "_" + strconv.Itoa(id)
}
