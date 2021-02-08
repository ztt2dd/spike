package kafka

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
	"spikeKill/models"
	"spikeKill/pkg/setting"
	"spikeKill/pkg/snowflake"
)

func Consumer(topic string) {
	log.Println("[Consumer]consumer start")
	consumer, err := sarama.NewConsumer([]string{setting.KafkaSetting.Host}, nil)
	if err != nil {
		log.Println("kafka consumer conn err:", err)
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Println("kafka consumer close err:", err)
		}
	}()

	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Println("kafka consumer partition err:", err)
	}

	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Println("kafka consumer partition close err:", err)
		}
	}()

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	consumed := 0
ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			var order *models.Orders
			json.Unmarshal(msg.Value, &order)
			err := DeductionLocalStock(order.ProductId, order.UserId)
			if err != nil {
				log.Println("[Consumer]consume err:", err)
			}
			log.Printf("Consumed message offset %d\n", msg.Offset)
			consumed++
		case <-signals:
			break ConsumerLoop
		}
	}

	log.Printf("Consumed: %d\n", consumed)
}

func DeductionLocalStock(productId int, userId int) error {
	orderSn := snowflake.GetSnowflakeId()
	order := &models.Orders{
		ProductId: productId,
		UserId:    userId,
		OrderSn:   orderSn,
	}
	return models.CreateLocalOrder(order)
}
