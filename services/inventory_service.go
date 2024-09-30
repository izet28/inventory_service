package services

import (
	"inventory_service/models"
	"inventory_service/repository"
	"inventory_service/utils"
	"net/http"
)

// InventoryService mendefinisikan logika bisnis untuk inventory
type InventoryService interface {
	GetInventory(productID uint) (*models.Inventory, *utils.AppError)
	UpdateStock(productID uint, newStock int) *utils.AppError
	ReplenishStock(productID uint, additionalStock int) *utils.AppError
	CreateInventory(inventory *models.Inventory) (*models.Inventory, *utils.AppError)
}

// inventoryService adalah implementasi dari InventoryService
type inventoryService struct {
	repository repository.InventoryRepository
}

// NewInventoryService membuat instance baru dari InventoryService
func NewInventoryService(repository repository.InventoryRepository) InventoryService {
	return &inventoryService{repository}
}

func (s *inventoryService) GetInventory(productID uint) (*models.Inventory, *utils.AppError) {
	inventory, err := s.repository.FindByProductID(productID)
	if err != nil {
		return nil, utils.NewAppError(http.StatusNotFound, "INVENTORY_NOT_FOUND", "Stok tidak ditemukan", err)
	}
	return inventory, nil
}

func (s *inventoryService) UpdateStock(productID uint, newStock int) *utils.AppError {
	err := s.repository.UpdateStock(productID, newStock)
	if err != nil {
		return utils.NewAppError(http.StatusInternalServerError, "INVENTORY_UPDATE_FAILED", "Gagal memperbarui stok", err)
	}
	return nil
}

func (s *inventoryService) ReplenishStock(productID uint, additionalStock int) *utils.AppError {
	err := s.repository.ReplenishStock(productID, additionalStock)
	if err != nil {
		return utils.NewAppError(http.StatusInternalServerError, "STOCK_REPLENISH_FAILED", "Gagal menambah stok", err)
	}
	return nil
}

func (s *inventoryService) CreateInventory(inventory *models.Inventory) (*models.Inventory, *utils.AppError) {
	createdInventory, err := s.repository.CreateInventory(inventory)
	if err != nil {
		return nil, utils.NewAppError(http.StatusInternalServerError, "INVENTORY_CREATION_FAILED", "Gagal membuat inventory", err)
	}
	return createdInventory, nil
}
