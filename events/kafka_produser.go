package events

import (
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
)

type KafkaProducer struct {
	producer sarama.SyncProducer
}

func NewKafkaProducer(brokers []string) (*KafkaProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &KafkaProducer{producer: producer}, nil
}

func (p *KafkaProducer) PublishProductRequest(topic string, productName string) error {
	requestData := map[string]string{"product_name": productName}
	productData, err := json.Marshal(requestData)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(productName), // Menggunakan nama produk sebagai key
		Value: sarama.ByteEncoder(productData),
	}

	partition, offset, err := p.producer.SendMessage(msg)
	if err != nil {
		return err
	}

	log.Printf("Request untuk produk %s berhasil dikirim ke Kafka. Partition: %d, Offset: %d\n", productName, partition, offset)
	return nil
}
