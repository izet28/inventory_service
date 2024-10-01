package repository

import (
	"inventory_service/models"

	"gorm.io/gorm"
)

type WarehouseRepository interface {
	CreateWarehouse(warehouse models.Warehouse) (models.Warehouse, error)
}

type warehouseRepository struct {
	db *gorm.DB
}

func NewWarehouseRepository(db *gorm.DB) WarehouseRepository {
	return &warehouseRepository{db}
}

func (r *warehouseRepository) CreateWarehouse(warehouse models.Warehouse) (models.Warehouse, error) {
	err := r.db.Create(&warehouse).Error
	if err != nil {
		return models.Warehouse{}, err
	}

	return warehouse, nil
}
