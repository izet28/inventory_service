package services

import (
	"inventory_service/models"
	"inventory_service/repository"

	"github.com/sirupsen/logrus"
)

type ProductWarehouseService interface {
	AddStockToWarehouse(productID, warehouseID uint, stock int) error
	UpdateStock(productID, warehouseID uint, stock int) error
	GetStockByProductAndWarehouse(productID, warehouseID uint) (*models.ProductWarehouse, error)
	GetAllStocksForProduct(productID uint) ([]models.ProductWarehouse, error)
}

type productWarehouseService struct {
	repository repository.ProductWarehouseRepository
	logger     *logrus.Logger
}

func NewProductWarehouseService(repo repository.ProductWarehouseRepository, logger *logrus.Logger) ProductWarehouseService {
	return &productWarehouseService{
		repository: repo,
		logger:     logger,
	}
}

// AddStockToWarehouse menambahkan stok produk ke warehouse
func (s *productWarehouseService) AddStockToWarehouse(productID, warehouseID uint, stock int) error {
	s.logger.Infof("Menambahkan stok produk %d ke warehouse %d dengan stok %d", productID, warehouseID, stock)

	// Cek apakah relasi product-warehouse sudah ada
	pw, err := s.repository.FindByProductAndWarehouse(productID, warehouseID)
	if err != nil {
		// Jika tidak ada, buat relasi baru
		productWarehouse := models.ProductWarehouse{
			ProductID:   productID,
			WarehouseID: warehouseID,
			Stock:       stock,
		}
		return s.repository.Create(&productWarehouse)
	}

	// Jika sudah ada, tambahkan stok
	pw.Stock += stock
	return s.repository.UpdateStock(productID, warehouseID, pw.Stock)
}

// UpdateStock memperbarui stok produk di warehouse
func (s *productWarehouseService) UpdateStock(productID, warehouseID uint, stock int) error {
	s.logger.Infof("Memperbarui stok produk %d di warehouse %d menjadi %d", productID, warehouseID, stock)
	return s.repository.UpdateStock(productID, warehouseID, stock)
}

// GetStockByProductAndWarehouse mendapatkan stok produk di warehouse tertentu
func (s *productWarehouseService) GetStockByProductAndWarehouse(productID, warehouseID uint) (*models.ProductWarehouse, error) {
	s.logger.Infof("Mengambil stok produk %d di warehouse %d", productID, warehouseID)
	return s.repository.FindByProductAndWarehouse(productID, warehouseID)
}

// GetAllStocksForProduct mendapatkan semua stok produk di semua warehouse
func (s *productWarehouseService) GetAllStocksForProduct(productID uint) ([]models.ProductWarehouse, error) {
	s.logger.Infof("Mengambil semua stok untuk produk %d di semua warehouse", productID)
	return s.repository.FindAllByProduct(productID)
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
