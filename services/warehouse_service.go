package services

import (
	"inventory_service/models"
	"inventory_service/repository"
	"inventory_service/utils"
	"net/http"
)

type WarehouseService interface {
	CreateWarehouse(warehouse *models.Warehouse) (*models.Warehouse, *utils.AppError)
}

type warehouseService struct {
	repository repository.WarehouseRepository
}

func NewWarehouseService(repository repository.WarehouseRepository) warehouseService {
	return warehouseService{repository}
}

func (s *warehouseService) CreateWarehouse(warehouse *models.Warehouse) (*models.Warehouse, *utils.AppError) {
	createdWarehouse, err := s.repository.CreateWarehouse(*warehouse)
	if err != nil {
		return nil, utils.NewAppError(http.StatusInternalServerError, "WAREHOUSE_CREATION_FAILED", "Gagal membuat warehouse", err)
	}

	return &createdWarehouse, nil
}
