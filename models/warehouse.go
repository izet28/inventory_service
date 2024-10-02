package models

import "gorm.io/gorm"

// Warehouse mewakili struktur tabel warehouse di database
type Warehouse struct {
	gorm.Model
	Name     string `json:"name" validate:"required,min=3,max=30"  `    // Nama warehouse dengan validasi
	Location string `json:"location" validate:"required,min=3,max=255"` // Lokasi warehouse dengan validasi
}
