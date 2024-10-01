package models

type Warehouse struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

type WarehouseStock struct {
	ID          uint `gorm:"primaryKey"`
	ProductID   uint ` json:"product_id"`
	WarehouseID uint `json:"warehouse_id"`
	Stock       int  `json:"stock"`
}
