package models

import "gorm.io/gorm"

// Inventory mewakili tabel inventory di database
type Inventory struct {
	gorm.Model
	ProductID uint `json:"product_id"` // ID produk
	Stock     int  `json:"stock"`      // Jumlah stok yang tersedia
}
