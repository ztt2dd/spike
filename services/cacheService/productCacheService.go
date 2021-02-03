package cacheService

import (
	"spikeKill/pkg/e"
	"strconv"
)

func GetProductKey(id int) string {
	return e.CACHE_PRODUCT + "_" + strconv.Itoa(id)
}
