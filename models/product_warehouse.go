package models

import "gorm.io/gorm"

// ProductWarehouse menyimpan stok dari produk di setiap warehouse
type ProductWarehouse struct {
	gorm.Model
	ProductID   uint `json:"product_id"`
	WarehouseID uint `json:"warehouse_id"`
	Stock       int  `json:"stock" gorm:"type:int;not null"` // Stok produk di warehouse tertentu
}
