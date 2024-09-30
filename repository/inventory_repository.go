package repository

import (
	"inventory_service/models"

	"gorm.io/gorm"
)

// InventoryRepository menyediakan antarmuka untuk operasi database
type InventoryRepository interface {
	FindByProductID(productID uint) (*models.Inventory, error)
	UpdateStock(productID uint, newStock int) error
	ReplenishStock(productID uint, additionalStock int) error
	CreateInventory(inventory *models.Inventory) (*models.Inventory, error)
}

// inventoryRepository adalah implementasi dari InventoryRepository
type inventoryRepository struct {
	db *gorm.DB
}

// NewInventoryRepository membuat instance baru dari InventoryRepository
func NewInventoryRepository(db *gorm.DB) InventoryRepository {
	return &inventoryRepository{db}
}

func (r *inventoryRepository) FindByProductID(productID uint) (*models.Inventory, error) {
	var inventory models.Inventory
	err := r.db.Where("product_id = ?", productID).First(&inventory).Error
	if err != nil {
		return nil, err
	}
	return &inventory, nil
}

func (r *inventoryRepository) UpdateStock(productID uint, newStock int) error {
	return r.db.Model(&models.Inventory{}).Where("product_id = ?", productID).Update("stock", newStock).Error
}

func (r *inventoryRepository) ReplenishStock(productID uint, additionalStock int) error {
	return r.db.Model(&models.Inventory{}).Where("product_id = ?", productID).UpdateColumn("stock", gorm.Expr("stock + ?", additionalStock)).Error
}

func (r *inventoryRepository) CreateInventory(inventory *models.Inventory) (*models.Inventory, error) {
	err := r.db.Create(inventory).Error
	if err != nil {
		return nil, err
	}
	return inventory, nil
}
