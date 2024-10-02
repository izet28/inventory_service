package repository

import (
	"inventory_service/models"

	"gorm.io/gorm"
)

// WarehouseRepository adalah interface untuk operasi CRUD pada Warehouse
type WarehouseRepository interface {
	Create(warehouse *models.Warehouse) error
	FindByID(id uint) (*models.Warehouse, error)
	FindAll() ([]models.Warehouse, error)
	Update(warehouse *models.Warehouse) error
	Delete(id uint) error
}

type warehouseRepository struct {
	db *gorm.DB
}

// NewWarehouseRepository membuat instance baru dari warehouseRepository
func NewWarehouseRepository(db *gorm.DB) WarehouseRepository {
	return &warehouseRepository{db}
}

// Create menyimpan warehouse baru ke database
func (r *warehouseRepository) Create(warehouse *models.Warehouse) error {
	return r.db.Create(warehouse).Error
}

// FindByID mencari warehouse berdasarkan ID
func (r *warehouseRepository) FindByID(id uint) (*models.Warehouse, error) {
	var warehouse models.Warehouse
	err := r.db.First(&warehouse, id).Error
	if err != nil {
		return nil, err
	}
	return &warehouse, nil
}

// FindAll mengambil semua warehouse dari database
func (r *warehouseRepository) FindAll() ([]models.Warehouse, error) {
	var warehouses []models.Warehouse
	err := r.db.Find(&warehouses).Error
	if err != nil {
		return nil, err
	}
	return warehouses, nil
}

// Update memperbarui warehouse di database
func (r *warehouseRepository) Update(warehouse *models.Warehouse) error {

	return r.db.Model(warehouse).Omit("CreatedAt").Save(warehouse).Error
}

// Delete menghapus warehouse dari database berdasarkan ID
func (r *warehouseRepository) Delete(id uint) error {
	return r.db.Delete(&models.Warehouse{}, id).Error
}
