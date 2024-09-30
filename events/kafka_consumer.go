package events

import (
	"encoding/json"
	"inventory_service/models"
	"inventory_service/services"
	"log"

	"github.com/IBM/sarama"
)

func ConsumeProductCreatedEvents(inventoryService services.InventoryService) {
	config := sarama.NewConfig()
	brokers := []string{"localhost:9092"} // Sesuaikan dengan konfigurasi Kafka
	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		log.Fatalf("Gagal terhubung ke Kafka: %v", err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("product-events", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Gagal membuat partition consumer: %v", err)
	}
	defer partitionConsumer.Close()

	for message := range partitionConsumer.Messages() {
		var product models.Product
		err := json.Unmarshal(message.Value, &product)
		if err != nil {
			log.Printf("Error decoding message: %v", err)
			continue
		}

		log.Printf("Menerima event produk: %+v", product)

		// Update stok inventory untuk produk yang baru dibuat
		inventory := models.Inventory{
			ProductID: product.ID,
			Stock:     product.Stock, // Asumsikan Product memiliki InitialStock
		}

		_, appErr := inventoryService.CreateInventory(&inventory)
		if appErr != nil {
			log.Printf("Gagal membuat inventory untuk produk %v: %v", product.ID, appErr)
		}
	}
}
