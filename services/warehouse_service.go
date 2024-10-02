package services

import (
	"errors"
	"inventory_service/models"
	"inventory_service/repository"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

// WarehouseService mengelola operasi warehouse
type WarehouseService interface {
	CreateWarehouse(warehouse *models.Warehouse) (*models.Warehouse, error)
	GetWarehouseByID(id uint) (*models.Warehouse, error)
	GetAllWarehouses() ([]models.Warehouse, error)
	UpdateWarehouse(warehouse *models.Warehouse) (*models.Warehouse, error)
	DeleteWarehouse(id uint) error
}

type warehouseService struct {
	repository repository.WarehouseRepository
	validate   *validator.Validate
	logger     *logrus.Logger
}

// NewWarehouseService membuat instance baru dari warehouseService
func NewWarehouseService(repo repository.WarehouseRepository, logger *logrus.Logger) WarehouseService {
	return &warehouseService{
		repository: repo,
		validate:   validator.New(),
		logger:     logger,
	}
}

// CreateWarehouse membuat warehouse baru
func (s *warehouseService) CreateWarehouse(warehouse *models.Warehouse) (*models.Warehouse, error) {
	s.logger.Infof("Mencoba membuat warehouse: %+v", warehouse)

	// Validasi input
	err := s.validate.Struct(warehouse)
	if err != nil {
		s.logger.Error("Validasi gagal: ", err)
		return nil, errors.New("input tidak valid")
	}

	// Simpan ke database
	err = s.repository.Create(warehouse)
	if err != nil {
		s.logger.Error("Gagal menyimpan warehouse: ", err)
		return nil, err
	}

	s.logger.Infof("Warehouse berhasil dibuat: %+v", warehouse)
	return warehouse, nil
}

// GetWarehouseByID mengambil warehouse berdasarkan ID
func (s *warehouseService) GetWarehouseByID(id uint) (*models.Warehouse, error) {
	s.logger.Infof("Mencari warehouse dengan ID: %d", id)

	warehouse, err := s.repository.FindByID(id)
	if err != nil {
		s.logger.Error("Warehouse tidak ditemukan: ", err)
		return nil, err
	}

	s.logger.Infof("Warehouse ditemukan: %+v", warehouse)
	return warehouse, nil
}

// GetAllWarehouses mengambil semua warehouse
func (s *warehouseService) GetAllWarehouses() ([]models.Warehouse, error) {
	s.logger.Info("Mengambil semua warehouse")

	warehouses, err := s.repository.FindAll()
	if err != nil {
		s.logger.Error("Gagal mengambil semua warehouse: ", err)
		return nil, err
	}

	return warehouses, nil
}

// UpdateWarehouse memperbarui warehouse yang ada
func (s *warehouseService) UpdateWarehouse(warehouse *models.Warehouse) (*models.Warehouse, error) {
	s.logger.Infof("Mencoba memperbarui warehouse dengan ID: %d", warehouse.ID)

	// Validasi input
	err := s.validate.Struct(warehouse)
	if err != nil {
		s.logger.Error("Validasi gagal: ", err)
		return nil, errors.New("input tidak valid")
	}

	err = s.repository.Update(warehouse)
	if err != nil {
		s.logger.Error("Gagal memperbarui warehouse: ", err)
		return nil, err
	}

	s.logger.Infof("Warehouse berhasil diperbarui: %+v", warehouse)
	return warehouse, nil
}

// DeleteWarehouse menghapus warehouse berdasarkan ID
func (s *warehouseService) DeleteWarehouse(id uint) error {
	s.logger.Infof("Menghapus warehouse dengan ID: %d", id)

	err := s.repository.Delete(id)
	if err != nil {
		s.logger.Error("Gagal menghapus warehouse: ", err)
		return err
	}

	s.logger.Infof("Warehouse dengan ID %d berhasil dihapus", id)
	return nil
}

func (s *productWarehouseService) AddStockByProductName(productName string, warehouseID uint, stock int) error {
	// Kirimkan request ke Kafka untuk meminta Product ID berdasarkan nama produk
	err := s.kafkaProducer.PublishProductRequest("product-request-topic", productName)
	if err != nil {
		s.logger.Errorf("Gagal mengirim request ke Kafka untuk produk %s: %v", productName, err)
		return err
	}

	s.logger.Infof("Request untuk produk %s telah dikirim ke Kafka, menunggu Product ID", productName)

	// Simpan stok setelah menerima Product ID (akan dijelaskan di bagian konsumer Kafka)
	// ...
	return nil
}
