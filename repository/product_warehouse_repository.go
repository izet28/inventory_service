package repository

import (
	"inventory_service/models"

	"gorm.io/gorm"
)

type ProductWarehouseRepository interface {
	Create(productWarehouse *models.ProductWarehouse) error
	FindByProductAndWarehouse(productID, warehouseID uint) (*models.ProductWarehouse, error)
	UpdateStock(productID, warehouseID uint, stock int) error
	FindAllByProduct(productID uint) ([]models.ProductWarehouse, error)
}

type productWarehouseRepository struct {
	db *gorm.DB
}

func NewProductWarehouseRepository(db *gorm.DB) ProductWarehouseRepository {
	return &productWarehouseRepository{db: db}
}

func (r *productWarehouseRepository) Create(productWarehouse *models.ProductWarehouse) error {
	return r.db.Create(productWarehouse).Error
}

func (r *productWarehouseRepository) FindByProductAndWarehouse(productID, warehouseID uint) (*models.ProductWarehouse, error) {
	var productWarehouse models.ProductWarehouse
	err := r.db.Where("product_id = ? AND warehouse_id = ?", productID, warehouseID).First(&productWarehouse).Error
	if err != nil {
		return nil, err
	}
	return &productWarehouse, nil
}

func (r *productWarehouseRepository) UpdateStock(productID, warehouseID uint, stock int) error {
	return r.db.Model(&models.ProductWarehouse{}).
		Where("product_id = ? AND warehouse_id = ?", productID, warehouseID).
		Update("stock", stock).Error
}

func (r *productWarehouseRepository) FindAllByProduct(productID uint) ([]models.ProductWarehouse, error) {
	var productWarehouses []models.ProductWarehouse
	err := r.db.Where("product_id = ?", productID).Find(&productWarehouses).Error
	if err != nil {
		return nil, err
	}
	return productWarehouses, nil
}
