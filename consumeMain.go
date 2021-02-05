package main

import (
	"spikeKill/models"
	"spikeKill/pkg/e"
	"spikeKill/pkg/kafka"
	"spikeKill/pkg/redis"
	"spikeKill/pkg/setting"
	"spikeKill/pkg/util"
)

func init() {
	setting.Setup()
	models.Setup()
	util.Setup()
	redis.Setup()
}

func consume() {
	kafka.Consumer(e.ORDER_TOPIC)
}

func main() {
	consume()
}
