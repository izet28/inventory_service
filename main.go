package main

import (
	"inventory_service/config"
	"inventory_service/controllers"
	"inventory_service/repository"
	"inventory_service/services"
	"inventory_service/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Inisialisasi database
	db := config.InitDB()
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	// Inisialisasi logger
	logger := utils.InitLogger()

	// Inisialisasi repository dan service
	warehouseRepo := repository.NewWarehouseRepository(db)
	productWarehouseRepo := repository.NewProductWarehouseRepository(db)
	warehouseService := services.NewWarehouseService(warehouseRepo, logger)
	productWarehouseService := services.NewProductWarehouseService(productWarehouseRepo, logger)
	warehouseController := controllers.NewWarehouseController(warehouseService, logger)
	productWarehouseController := controllers.NewProductWarehouseController(productWarehouseService)

	// Inisialisasi Kafka Consumer (jika diperlukan)
	// kafkaConsumer := events.NewKafkaConsumer(productWarehouseService)
	// go kafkaConsumer.ConsumeProductEvents([]string{"localhost:9092"}, "product-topic")

	// Buat router
	r := mux.NewRouter()

	// Route CRUD warehouse
	r.HandleFunc("/warehouses", warehouseController.CreateWarehouse).Methods("POST")
	r.HandleFunc("/warehouses/{id}", warehouseController.GetWarehouseByID).Methods("GET")
	r.HandleFunc("/warehouses", warehouseController.GetAllWarehouses).Methods("GET")
	r.HandleFunc("/warehouses/{id}", warehouseController.UpdateWarehouse).Methods("PUT")
	r.HandleFunc("/warehouses/{id}", warehouseController.DeleteWarehouse).Methods("DELETE")

	// Route untuk mengelola stok produk di warehouse
	r.HandleFunc("/products/{product_id}/warehouses/{warehouse_id}/stock", productWarehouseController.AddStockToWarehouse).Methods("POST")
	r.HandleFunc("/products/{product_id}/warehouses/{warehouse_id}/stock", productWarehouseController.UpdateStock).Methods("PUT")
	r.HandleFunc("/products/{product_id}/warehouses/{warehouse_id}/stock", productWarehouseController.GetStock).Methods("GET")

	log.Println("Server berjalan pada port 8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
