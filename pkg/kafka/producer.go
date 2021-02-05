package kafka

import (
	"github.com/Shopify/sarama"
	"log"
	"spikeKill/pkg/setting"
)

func Producer(topic string, data string) error {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// config.Producer.Return.Successes = true // 只有同步消息才设置为true
	producer, err := sarama.NewAsyncProducer([]string{setting.KafkaSetting.Host}, config)
	if err != nil {
		return err
	}
	defer producer.AsyncClose()

	producer.Input() <- &sarama.ProducerMessage{Topic: topic, Key: nil, Value: sarama.StringEncoder(data)}
	select {

	// 同步消息才会接收到success
	/*case msg := <-producer.Successes():
	log.Printf("Produced message successes: [%s]\n", msg.Value)*/
	case err := <-producer.Errors():
		log.Println("Produced message failure: ", err)
	default:
		log.Println("Produced message default")
	}
	if err != nil {
		return err
	}
	return nil
}
