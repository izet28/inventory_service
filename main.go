package main

import (
	"inventory_service/config"
	"inventory_service/controllers"
	"inventory_service/events"
	"inventory_service/repository"
	"inventory_service/services"
	"inventory_service/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	logger := utils.InitLogger()

	db := config.InitDB()
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	inventoryRepository := repository.NewInventoryRepository(db)
	inventoryService := services.NewInventoryService(inventoryRepository)
	inventoryController := controllers.NewInventoryController(inventoryService, logger)

	go events.ConsumeProductCreatedEvents(inventoryService)

	r := mux.NewRouter()

	r.HandleFunc("/inventory/{product_id}", inventoryController.GetInventory).Methods("GET")
	r.HandleFunc("/inventory/{product_id}", inventoryController.UpdateStock).Methods("PUT")
	r.HandleFunc("/inventory/{product_id}/replenish", inventoryController.ReplenishStock).Methods("PUT")

	log.Println("Inventory Service berjalan pada port 8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
